import { NavLink } from 'react-router-dom'
import { cn } from '@/lib/utils'
import {
  MessageSquare,
  FolderKanban,
  CheckSquare,
  GitPullRequest,
  Users,
  Bot,
} from 'lucide-react'

const navItems = [
  { to: '/', icon: Bot, label: '电子助理' },
  { to: '/projects', icon: FolderKanban, label: '项目管理' },
  { to: '/tasks', icon: CheckSquare, label: '智能任务' },
  { to: '/review', icon: GitPullRequest, label: '代码审查' },
  { to: '/team', icon: Users, label: '团队协作' },
]

export function Sidebar() {
  return (
    <aside className="flex w-64 flex-col border-r bg-card">
      <div className="flex h-14 items-center border-b px-4">
        <MessageSquare className="mr-2 h-6 w-6 text-primary" />
        <span className="text-lg font-semibold">AI 研发助手</span>
      </div>
      <nav className="flex-1 overflow-y-auto p-4">
        <ul className="space-y-1">
          {navItems.map((item) => (
            <li key={item.to}>
              <NavLink
                to={item.to}
                className={({ isActive }) =>
                  cn(
                    'flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors',
                    isActive
                      ? 'bg-primary text-primary-foreground'
                      : 'text-muted-foreground hover:bg-accent hover:text-accent-foreground'
                  )
                }
              >
                <item.icon className="h-5 w-5" />
                {item.label}
              </NavLink>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  )
}
