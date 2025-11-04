import type { Comment } from '@/types'

export interface CommentNode extends Comment {
  children: CommentNode[]
}

export function buildTree(comments: Comment[]): CommentNode[] {
  const commentsMap: { [key: string]: CommentNode } = {}
  const tree: CommentNode[] = []

  for (const comment of comments) {
    const commentNode: CommentNode = { ...comment, children: [] }
    commentsMap[commentNode.id] = commentNode
  }

  for (const commentID in commentsMap) {
    const commentNode = commentsMap[commentID]
    if (!commentNode) {
      continue
    }

    if (commentNode.parent_id) {
      const parent = commentsMap[commentNode.parent_id]
      if (parent) {
        parent.children.push(commentNode)
      } else {
        tree.push(commentNode)
      }
    } else {
      tree.push(commentNode)
    }
  }

  return tree
}
