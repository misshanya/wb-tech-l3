<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { Comment } from './types'
import { deleteComment, searchComments, createComment, getComments } from './services/api'
import CommentItem from './components/CommentItem.vue'
import Input from './components/ui/input/Input.vue'
import { buildTree } from './lib/tree'
import type { CommentNode } from './lib/tree'

const isLoading = ref(false)

const error = ref<string | null>(null)

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
    isLoading.value = true
    error.value = null

    try {
      const result = await searchComments(newQuery)
      comments.value = result
    } catch (e) {
      error.value = (e as Error).message
      console.log(`failed to search comments: ${e}`)
    } finally {
      isLoading.value = false
    }
  }, 500)
})

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
  <div class="max-w-3xl mx-auto mt-8">
    <Input v-model="searchQuery" placeholder="Поиск..." />
    <div class="comments">
      <div v-if="isLoading">Загрузка...</div>
      <div v-else-if="error">{{ error }}</div>
      <div v-else>
        <CommentItem
          class="mt-4"
          v-for="comment in commentsTree"
          :key="comment.id"
          :comment="comment"
          @delete="handleCommentDelete"
          @reply="handleCommentCreate"
          @expand="handleCommentExpand"
        />
      </div>
    </div>
  </div>
</template>
