export interface Comment {
  id: string
  content: string
  created_at: string
  parent_id: string | null
}

export interface CommentsApiResponse {
  comments: Comment[]
}
