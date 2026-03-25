# AI 研发助手 - 系统架构设计

## 1. 系统概述

### 1.1 架构设计原则
- **前后端分离**：前端负责 UI 交互，后端负责业务逻辑
- **服务化设计**：各模块独立部署，通过 API 通信
- **可扩展性**：支持水平扩展，适应不同规模团队
- **安全性**：完善的认证授权机制，保护代码安全

### 1.2 系统架构图

```
┌─────────────────────────────────────────────────────────────────┐
│                         客户端层                                  │
├──────────┬──────────┬──────────┬──────────┬──────────┬───────────┤
│   Web    │  桌面端  │  移动端  │  CLI     │  VSCode  │  GitBot   │
│  (React) │(Electron)│(Capacitor)│ (Go)    │  Plugin  │ (Webhook) │
└──────────┴──────────┴──────────┴──────────┴──────────┴───────────┘
                               │
                               │ HTTPS/WSS
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                         网关层 (Gateway)                           │
│                   Nginx / Envoy (负载均衡 + SSL)                 │
└─────────────────────────────────────────────────────────────────┘
                               │
          ┌────────────────────┼────────────────────┐
          │                    │                    │
          ▼                    ▼                    ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│   Web 前端服务   │  │   API 服务      │  │  WebSocket 服务  │
│   (Frontend)    │  │   (Backend)     │  │   (Realtime)    │
│   Port: 3000    │  │   Port: 8080   │  │   Port: 8081    │
└─────────────────┘  └─────────────────┘  └─────────────────┘
          │                    │                    │
          └────────────────────┼────────────────────┘
                               │
          ┌────────────────────┼────────────────────┐
          │                    │                    │
          ▼                    ▼                    ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│   PostgreSQL    │  │     Redis       │  │   文件存储       │
│   (主数据库)     │  │   (缓存/会话)    │  │   (S3/MinIO)    │
└─────────────────┘  └─────────────────┘  └─────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                       AI 模型层 (AI Gateway)                      │
├──────────┬──────────┬──────────┬──────────┬──────────┬───────────┤
│  OpenAI  │  Claude  │  通义千问  │   智谱   │  硅基流动  │  Ollama   │
│   API    │   API    │    API   │   API    │    API   │  (本地)   │
└──────────┴──────────┴──────────┴──────────┴──────────┴───────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                     开发环境层 (Dev Environments)                 │
├──────────┬──────────┬──────────┬──────────┬──────────┬───────────┤
│  Code    │  Terminal │  File    │  Port    │   Port   │   Git     │
│  Server  │  Manager │  Manager │  Forward │  Manager │  Sync     │
│ (VSCode) │          │          │          │          │           │
└──────────┴──────────┴──────────┴──────────┴──────────┴───────────┘
```

---

## 2. 前端架构 (Frontend)

### 2.1 技术选型
- **框架**：React 19
- **语言**：TypeScript 5.x
- **构建工具**：Vite 5.x
- **样式**：Tailwind CSS 3.x
- **状态管理**：Zustand / React Query
- **路由**：React Router 6.x
- **UI 组件**：shadcn/ui + Radix UI

### 2.2 前端模块结构

```
frontend/
├── public/                 # 静态资源
├── src/
│   ├── components/        # 公共组件
│   │   ├── ui/            # UI 基础组件
│   │   ├── layout/        # 布局组件
│   │   └── ide/           # IDE 相关组件
│   ├── features/          # 功能模块
│   │   ├── tasks/         # 智能任务
│   │   ├── projects/      # 项目管理
│   │   ├── ide/           # 在线开发环境
│   │   ├── review/        # 代码审查
│   │   └── team/          # 团队协作
│   ├── hooks/             # 自定义 Hooks
│   ├── services/          # API 服务
│   ├── stores/            # 状态管理
│   ├── types/             # TypeScript 类型
│   └── utils/             # 工具函数
├── package.json
└── vite.config.ts
```

### 2.3 前端核心功能

#### 2.3.1 智能任务模块 (tasks)
- 任务创建与编辑
- 任务状态管理
- AI 对话交互
- 代码预览与编辑

#### 2.3.2 在线 IDE 模块 (ide)
- Monaco Editor 集成
- 终端模拟器
- 文件浏览器
- 多标签页管理

#### 2.3.3 项目管理模块 (projects)
- 项目列表与详情
- 需求文档管理
- 任务看板

---

## 3. 后端架构 (Backend)

### 3.1 技术选型
- **语言**：Go 1.21+
- **框架**：Echo v4
- **ORM**：Ent
- **数据库**：PostgreSQL 16
- **缓存**：Redis 7
- **文件存储**：S3 / MinIO
- **AI 网关**：自定义实现

### 3.2 后端模块结构

```
backend/
├── cmd/
│   └── server/            # 主服务入口
├── internal/
│   ├── api/               # API 处理层
│   │   ├── handlers/      # HTTP Handlers
│   │   ├── middleware/    # 中间件
│   │   └── router/        # 路由定义
│   ├── service/           # 业务逻辑层
│   │   ├── task/          # 任务服务
│   │   ├── project/       # 项目服务
│   │   ├── ai/            # AI 服务
│   │   ├── ide/           # IDE 服务
│   │   ├── review/        # 审查服务
│   │   └── team/          # 团队服务
│   ├── repository/        # 数据访问层
│   │   ├── ent/           # Ent ORM 模型
│   │   └── cache/         # Redis 缓存
│   ├── model/             # 数据模型
│   └── pkg/               # 公共包
├── pkg/
│   ├── ai/                # AI 模型客户端
│   ├── git/               # Git 操作库
│   └── docker/            # Docker 操作库
├── migrations/            # 数据库迁移
└── go.mod
```

### 3.3 数据库设计

#### 3.3.1 核心实体

**User（用户）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| email | string | 邮箱 |
| name | string | 名称 |
| avatar | string | 头像 URL |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

**Team（团队）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| name | string | 团队名称 |
| owner_id | uuid | 所有者 ID |
| created_at | timestamp | 创建时间 |

**Project（项目）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| name | string | 项目名称 |
| description | string | 项目描述 |
| team_id | uuid | 团队 ID |
| git_url | string | Git 仓库地址 |
| created_at | timestamp | 创建时间 |

**Task（任务）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| title | string | 任务标题 |
| description | text | 任务描述 |
| type | enum | 任务类型（dev/design/review） |
| status | enum | 任务状态 |
| project_id | uuid | 项目 ID |
| created_by | uuid | 创建者 ID |
| created_at | timestamp | 创建时间 |

**TaskMessage（任务消息）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| task_id | uuid | 任务 ID |
| role | enum | 角色（user/assistant） |
| content | text | 消息内容 |
| created_at | timestamp | 创建时间 |

---

## 4. AI 模型集成

### 4.1 AI 网关架构

```
                    ┌─────────────────┐
                    │   AI Gateway    │
                    │   (ai/gateway)  │
                    └────────┬────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│  OpenAI       │  │   Anthropic  │  │  国内模型      │
│  Adapter      │  │   Adapter    │  │  Adapter      │
└───────────────┘  └───────────────┘  └───────────────┘
```

### 4.2 支持的 AI 模型

| 模型 | Provider | 说明 |
|------|----------|------|
| GPT-4o | OpenAI | 主力和模型 |
| GPT-4o-mini | OpenAI | 轻量级模型 |
| Claude 3.5 Sonnet | Anthropic | 代码能力强 |
| Claude 3.5 Haiku | Anthropic | 轻量级模型 |
| Qwen-Max | 通义千问 | 阿里大模型 |
| GLM-4-Plus | 智谱 | 智谱大模型 |
| SiliconFlow | 硅基流动 | 聚合 API |

### 4.3 AI 服务接口

```go
// AI Provider 接口
type AIProvider interface {
    Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error)
    Models() []Model
}

// ChatRequest
type ChatRequest struct {
    Model       string
    Messages    []Message
    Temperature float64
    MaxTokens   int
}

// Message
type Message struct {
    Role    string // "user", "assistant", "system"
    Content string
}
```

---

## 5. 在线开发环境架构

### 5.1 开发环境组件

```
┌─────────────────────────────────────────────────────┐
│              Web IDE (Browser)                       │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐        │
│  │  Editor  │  │ Terminal │  │  File    │        │
│  │ (Monaco) │  │ (xterm.js)│  │ Browser  │        │
│  └──────────┘  └──────────┘  └──────────┘        │
└─────────────────────────────────────────────────────┘
                         │
                         │ WebSocket
                         ▼
┌─────────────────────────────────────────────────────┐
│            IDE Backend Service                       │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐        │
│  │  Code    │  │ Terminal  │  │  File    │        │
│  │  Server  │  │ Manager   │  │ Manager  │        │
│  └──────────┘  └──────────┘  └──────────┘        │
└─────────────────────────────────────────────────────┘
                         │
                         │ SSH / Docker
                         ▼
┌─────────────────────────────────────────────────────┐
│           Development Container                      │
│  ┌──────────────────┐  ┌──────────────────┐        │
│  │  VSCode Server   │  │   Ubuntu/Debian  │        │
│  │  (code-server)   │  │   + 开发工具链    │        │
│  └──────────────────┘  └──────────────────┘        │
└─────────────────────────────────────────────────────┘
```

### 5.2 开发环境功能

| 功能 | 实现方案 | 说明 |
|------|----------|------|
| 代码编辑 | Monaco Editor | 支持多语言语法高亮 |
| 终端 | xterm.js + WebSocket | 浏览器内终端 |
| 文件管理 | WebDAV / SFTP | 在线文件浏览编辑 |
| 端口转发 | WebSocket Proxy | 预览 Web 服务 |
| 协作编辑 | WebSocket + CRDT | 多人实时协作 |

---

## 6. 代码审查架构

### 6.1 Git 机器人架构

```
┌─────────────┐     Webhook      ┌─────────────┐
│   GitHub    │ ───────────────► │  Review     │
│   GitLab    │                  │  Bot        │
│   Gitee     │                  │  Service    │
└─────────────┘                  └──────┬──────┘
                                         │
                                         │ API
                                         ▼
                                  ┌─────────────┐
                                  │   AI        │
                                  │   Review    │
                                  │   Engine    │
                                  └─────────────┘
```

### 6.2 审查流程

1. **PR/MR 创建**：Git 平台触发 Webhook
2. **代码获取**：Bot 获取代码变更
3. **AI 分析**：调用 AI 模型分析代码
4. **意见生成**：生成审查意见
5. **评论提交**：Bot 在 PR/MR 下发表评论

---

## 7. 安全架构

### 7.1 认证授权

- **JWT Token**：无状态认证
- **Refresh Token**：Token 续期
- **OAuth 2.0**：第三方登录（GitHub/GitLab/Gitee）
- **RBAC**：基于角色的权限控制

### 7.2 数据安全

- **传输加密**：全站 HTTPS
- **存储加密**：敏感数据加密存储
- **代码安全**：代码不落地，本地处理后即销毁

---

## 8. 部署架构

### 8.1 Docker Compose 部署

```yaml
version: '3.8'
services:
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgresql://user:pass@db:5432/monkeycode
      - REDIS_URL=redis://redis:6379

  db:
    image: postgres:16
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:
```

### 8.2 Kubernetes 部署

支持 Helm Chart 部署，提供：
- Horizontal Pod Autoscaler
- Rolling Update
- ConfigMap / Secret 管理
