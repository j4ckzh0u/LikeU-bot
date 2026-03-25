import { GitPullRequest } from 'lucide-react'

export function ReviewPage() {
  return (
    <div className="flex h-full flex-col">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">代码审查</h1>
        <p className="text-muted-foreground">自动审查 PR/MR 代码变更</p>
      </div>
      <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed">
        <div className="text-center">
          <GitPullRequest className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-medium">审查列表</h3>
          <p className="text-sm text-muted-foreground">暂无审查任务</p>
        </div>
      </div>
    </div>
  )
}
