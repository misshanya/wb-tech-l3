<script setup lang="ts">
import { buildTree, type CommentNode } from '@/lib/tree'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardFooter } from '@/components/ui/card'
import { Textarea } from '@/components/ui/textarea'
import { computed, ref, Transition } from 'vue'

const props = defineProps<{
  comment: CommentNode
}>()

const isReplyFormVisible = ref<boolean>(false)
const newReplyText = ref<string>('')

const emit = defineEmits(['delete', 'reply', 'expand'])

function onDeleteClick() {
  emit('delete', props.comment.id)
}

function handleSubmitReply() {
  emit('reply', {
    parentId: props.comment.id,
    content: newReplyText.value,
  })
}

const isExpanded = ref(true)
const loaded = ref(false)

const hasChildren = computed(() => props.comment.children && props.comment.children.length > 0)

async function handleExpandClick() {
  if (loaded.value && hasChildren.value) {
    isExpanded.value = !isExpanded.value
  } else {
    emit('expand', props.comment.id)
    loaded.value = true
    isExpanded.value = true
  }
}
</script>

<template>
  <div>
    <Card class="comment-item">
      <CardContent>
        {{ comment.content }}
      </CardContent>
      <CardFooter class="flex flex-col items-start gap-4">
        <span class="text-xs text-muted-foreground">
          {{ new Date(comment.created_at).toLocaleString('ru-RU') }}
        </span>
        <div class="flex items-center gap-2">
          <Button variant="outline" size="sm" @click="isReplyFormVisible = !isReplyFormVisible"
            >Ответить</Button
          >
          <Button variant="destructive" size="sm" @click="onDeleteClick">Удалить</Button>

          <Button @click="handleExpandClick" variant="outline" size="sm">
            {{ loaded && isExpanded && hasChildren ? 'Свернуть' : 'Развернуть' }}
          </Button>
        </div>
      </CardFooter>

      <Transition name="fade">
        <div v-if="isReplyFormVisible" class="p-6 pt-0">
          <Textarea v-model="newReplyText" placeholder="Ответ..." class="mb-4" />

          <div class="flex justify-end gap-2">
            <Button @click="handleSubmitReply">Отправить</Button>
          </div>
        </div>
      </Transition>
    </Card>
    <div v-if="hasChildren && isExpanded" class="ml-4 border-l-2 border-slate-700">
      <CommentItem
        v-for="child in comment.children"
        :key="child.id"
        :comment="child"
        class="mt-4"
        @delete="$emit('delete', $event)"
        @reply="$emit('reply', $event)"
        @expand="$emit('expand', $event)"
      />
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.1s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
