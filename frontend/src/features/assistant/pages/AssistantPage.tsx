import { Bot } from 'lucide-react'

export function AssistantPage() {
  return (
    <div className="flex h-full flex-col">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">电子助理</h1>
        <p className="text-muted-foreground">与 AI 助手对话，完成开发任务</p>
      </div>
      <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed">
        <div className="text-center">
          <Bot className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-medium">电子助理</h3>
          <p className="text-sm text-muted-foreground">开始与 AI 助手对话</p>
        </div>
      </div>
    </div>
  )
}
