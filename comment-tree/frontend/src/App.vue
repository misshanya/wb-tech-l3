<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { Comment } from './types'
import { deleteComment, searchComments, createComment, getComments } from './services/api'
import CommentItem from './components/CommentItem.vue'
import NewThreadForm from './components/NewThreadForm.vue'
import Input from './components/ui/input/Input.vue'
import { buildTree } from './lib/tree'
import type { CommentNode } from './lib/tree'
import { useColorMode } from '@vueuse/core'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from './components/ui/dropdown-menu'
import { Icon } from '@iconify/vue'
import { Button } from '@/components/ui/button'
import { NumberField, NumberFieldContent } from './components/ui/number-field'
import Label from './components/ui/label/Label.vue'
import NumberFieldDecrement from './components/ui/number-field/NumberFieldDecrement.vue'
import NumberFieldInput from './components/ui/number-field/NumberFieldInput.vue'
import NumberFieldIncrement from './components/ui/number-field/NumberFieldIncrement.vue'

const mode = useColorMode()

const isLoading = ref(false)

const error = ref<string | null>(null)

const searchPage = ref(1)
const searchLimit = ref(10)
const hasMoreComments = ref(true)
const searchQuery = ref<string>('')
const comments = ref<Comment[]>([])

const commentsTree = computed<CommentNode[]>(() => {
  if (comments.value.length === 0) {
    return []
  }
  return buildTree(comments.value)
})

let debounceTimer: ReturnType<typeof setTimeout>

watch(searchQuery, (newQuery) => {
  clearTimeout(debounceTimer)

  debounceTimer = setTimeout(async () => {
    comments.value = []
    searchPage.value = 1
    hasMoreComments.value = true
    isLoading.value = true
    error.value = null

    try {
      const result = await searchComments(newQuery, searchLimit.value, searchPage.value)
      comments.value.push(...result)
      searchPage.value++
    } catch (e) {
      error.value = (e as Error).message
      console.log(`failed to search comments: ${e}`)
    } finally {
      isLoading.value = false
    }
  }, 500)
})

async function handleLoadMoreComments() {
  try {
    const result = await searchComments(searchQuery.value, searchLimit.value, searchPage.value)
    if (result.length === 0) {
      hasMoreComments.value = false
      return
    }
    comments.value.push(...result)
    searchPage.value++
  } catch (e) {
    error.value = (e as Error).message
    console.log(`failed to load more comments: ${e}`)
  }
}

async function handleCommentDelete(commentID: string) {
  comments.value = comments.value.filter((comment) => comment.id !== commentID)

  try {
    await deleteComment(commentID)
  } catch (e) {
    error.value = (e as Error).message
    console.log(`failed to delete comment: ${e}`)
  }
}

interface ReplyPayload {
  parentId: string | null
  content: string
}

async function handleCommentCreate(payload: ReplyPayload) {
  try {
    const newComment = await createComment(payload.parentId, payload.content)
    comments.value.push(newComment)
  } catch (e) {
    error.value = (e as Error).message
    console.log(`failed to create comment: ${e}`)
  }
}

async function handleCommentExpand(parentID: string) {
  try {
    const newLoadedComments = await getComments(parentID)

    const existingIds = new Set(comments.value.map((c) => c.id))
    const newComments = newLoadedComments.filter((c) => !existingIds.has(c.id))

    comments.value.push(...newComments)
  } catch (e) {
    error.value = (e as Error).message
    console.log(`failed to load expanded comments: ${e}`)
  }
}
</script>

<template>
  <div class="p-8">
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline">
            <Icon
              icon="radix-icons:moon"
              class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Icon
              icon="radix-icons:sun"
              class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Сменить тему</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem @click="mode = 'light'"> Светлая </DropdownMenuItem>
          <DropdownMenuItem @click="mode = 'dark'"> Темная </DropdownMenuItem>
          <DropdownMenuItem @click="mode = 'auto'"> Системная </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>

    <div class="max-w-3xl mx-auto mt-8">
      <NewThreadForm class="mb-6" @submit="handleCommentCreate" />

      <NumberField v-model="searchLimit" :min="0" class="mb-6">
        <Label>Лимит</Label>
        <NumberFieldContent>
          <NumberFieldDecrement />
          <NumberFieldInput />
          <NumberFieldIncrement />
        </NumberFieldContent>
      </NumberField>

      <Input v-model="searchQuery" placeholder="Поиск..." />

      <div class="comments">
        <div v-if="isLoading">Загрузка...</div>
        <div v-else-if="error">{{ error }}</div>
        <div v-else class="grid w-full">
          <CommentItem
            class="mt-4"
            v-for="comment in commentsTree"
            :key="comment.id"
            :comment="comment"
            @delete="handleCommentDelete"
            @reply="handleCommentCreate"
            @expand="handleCommentExpand"
          />

          <Button
            v-if="hasMoreComments && comments.length > 0"
            @click="handleLoadMoreComments"
            class="mt-8"
          >Загрузить ещё</Button
          >
        </div>
      </div>
    </div>
  </div>
</template>
