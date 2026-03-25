import { FolderKanban } from 'lucide-react'

export function ProjectsPage() {
  return (
    <div className="flex h-full flex-col">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">项目管理</h1>
        <p className="text-muted-foreground">管理您的项目需求和任务</p>
      </div>
      <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed">
        <div className="text-center">
          <FolderKanban className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-medium">项目列表</h3>
          <p className="text-sm text-muted-foreground">暂无项目</p>
        </div>
      </div>
    </div>
  )
}
