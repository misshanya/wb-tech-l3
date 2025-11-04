import type { Comment, CommentsApiResponse } from '@/types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function getComments(parentId: string): Promise<Comment[]> {
  let url = `${API_BASE_URL}/comment?parent=${parentId}`

  const response = await fetch(url)
  if (!response.ok) {
    throw new Error(`Failed to fetch comment: status ${response.status}`)
  }

  const data: CommentsApiResponse = await response.json()
  return data.comments
}

export async function searchComments(query: string): Promise<Comment[]> {
  let url = `${API_BASE_URL}/comment/search?q=${query}`

  const response = await fetch(url)
  if (!response.ok) {
    throw new Error(`Failed to search comments: status ${response.status}`)
  }

  const data: CommentsApiResponse = await response.json()
  return data.comments
}

export async function deleteComment(id: string) {
  let url = `${API_BASE_URL}/comment/${id}`

  const response = await fetch(url, {
    method: 'DELETE',
  })
  if (!response.ok) {
    throw new Error(`Failed to delete comment: status ${response.status}`)
  }
}

export async function createComment(parentID: string | null, content: string): Promise<Comment> {
  let url = `${API_BASE_URL}/comment`

  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ content, parent_id: parentID }),
  })

  if (!(response.status == 201)) {
    throw new Error(`Failed to create comment: status ${response.status}`)
  }

  return response.json()
}
