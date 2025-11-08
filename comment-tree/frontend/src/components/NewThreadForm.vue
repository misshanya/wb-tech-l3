<script setup lang="ts">
import { ref } from 'vue'
import Card from './ui/card/Card.vue'
import CardContent from './ui/card/CardContent.vue'
import CardDescription from './ui/card/CardDescription.vue'
import CardHeader from './ui/card/CardHeader.vue'
import CardTitle from './ui/card/CardTitle.vue'
import Textarea from './ui/textarea/Textarea.vue'
import CardFooter from './ui/card/CardFooter.vue'
import Button from './ui/button/Button.vue'

const content = ref('')

const emit = defineEmits(['submit'])

async function handleSubmitButton() {
  emit('submit', {
    content: content.value,
  })
}
</script>

<template>
  <div>
    <Card>
      <CardHeader>
        <CardTitle>Начать тред</CardTitle>
        <CardDescription>Создание нового корневого комментария</CardDescription>
      </CardHeader>
      <CardContent class="grid w-full gap-1.5">
        <Label for="comment-content">Текст комментария</Label>
        <Textarea id="comment-content" v-model="content" placeholder="Привет!" />
      </CardContent>
      <CardFooter class="grid w-full">
        <Button @click="handleSubmitButton">Отправить</Button>
      </CardFooter>
    </Card>
  </div>
</template>
