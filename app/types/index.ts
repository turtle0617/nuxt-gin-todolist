export enum TodoStatus {
  COMPLETED = 'completed',
  IDLE = 'idle',
}

export interface Todo {
  id: string
  status: TodoStatus
  title: string
}
