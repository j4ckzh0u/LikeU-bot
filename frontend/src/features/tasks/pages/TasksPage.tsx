import { CheckSquare } from 'lucide-react'

export function TasksPage() {
  return (
    <div className="flex h-full flex-col">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">智能任务</h1>
        <p className="text-muted-foreground">创建和管理开发任务</p>
      </div>
      <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed">
        <div className="text-center">
          <CheckSquare className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-medium">任务列表</h3>
          <p className="text-sm text-muted-foreground">暂无任务</p>
        </div>
      </div>
    </div>
  )
}
