import { Users } from 'lucide-react'

export function TeamPage() {
  return (
    <div className="flex h-full flex-col">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">团队协作</h1>
        <p className="text-muted-foreground">管理团队成员和资源</p>
      </div>
      <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed">
        <div className="text-center">
          <Users className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-medium">团队成员</h3>
          <p className="text-sm text-muted-foreground">暂无团队</p>
        </div>
      </div>
    </div>
  )
}
