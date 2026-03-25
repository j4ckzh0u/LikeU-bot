export interface User {
  id: string
  email: string
  name: string
  avatar?: string
  created_at: string
}

export interface Team {
  id: string
  name: string
  owner_id: string
  created_at: string
}

export interface Project {
  id: string
  name: string
  description: string
  team_id: string
  git_url?: string
  created_at: string
}

export interface Task {
  id: string
  title: string
  description: string
  type: 'dev' | 'design' | 'review'
  status: 'pending' | 'in_progress' | 'completed'
  project_id?: string
  created_by: string
  created_at: string
}

export interface Agent {
  id: string
  name: string
  type: 'dev' | 'review' | 'project' | 'research' | 'custom'
  config: AgentConfig
  workspace?: string
}

export interface AgentConfig {
  model: string
  temperature: number
  max_tokens: number
}

export interface Channel {
  id: string
  type: 'web' | 'telegram' | 'discord' | 'slack' | 'whatsapp'
  name: string
  config: Record<string, unknown>
  enabled: boolean
}

export interface Session {
  id: string
  agent_id: string
  channel: string
  account_id: string
  status: 'active' | 'closed'
  created_at: string
}

export interface Message {
  id: string
  role: 'user' | 'assistant' | 'system'
  content: string
  parent_id?: string
  created_at: string
}

export interface Skill {
  id: string
  name: string
  version: string
  description: string
  enabled: boolean
  builtin: boolean
}

export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

export interface PaginatedResponse<T> {
  items: T[]
  total: number
  page: number
  page_size: number
}
