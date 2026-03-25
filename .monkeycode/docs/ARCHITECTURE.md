# AI 研发助手 - 系统架构设计

## 1. 系统概述

### 1.1 架构设计原则
- **前后端分离**：前端负责 UI 交互，后端负责业务逻辑
- **服务化设计**：各模块独立部署，通过 API 通信
- **可扩展性**：支持水平扩展，适应不同规模团队
- **安全性**：完善的认证授权机制，保护代码安全
- **本地优先**：数据本地处理，保护用户隐私

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
│                    Gateway 网关层                                  │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │              AI Assistant Gateway (电子助理中枢)           │    │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────┐ │    │
│  │  │  Channel │  │  Multi-  │  │  Skill   │  │ Session│ │    │
│  │  │  Manager │  │  Agent   │  │  Engine  │  │ Manager│ │    │
│  │  └──────────┘  └──────────┘  └──────────┘  └────────┘ │    │
│  └─────────────────────────────────────────────────────────┘    │
│                   Nginx / Envoy (负载均衡 + SSL)                  │
└─────────────────────────────────────────────────────────────────┘
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
           ┌───────────────────┼───────────────────┐
           │                   │                    │
           ▼                   ▼                    ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│  开发环境层      │  │   Canvas 层     │  │   语音服务层     │
│ (Dev Environ)   │  │  (Live Canvas)  │  │   (Voice)      │
└─────────────────┘  └─────────────────┘  └─────────────────┘
```

---

## 2. 电子助理模块 (AI Assistant)

### 2.1 模块概述

电子助理模块是系统的核心创新，融合 OpenClaw 的设计理念，提供：

- **多平台消息集成**：统一管理各种消息渠道
- **多 Agent 路由**：智能分发请求到专用 Agent
- **技能系统**：可扩展的能力体系
- **实时画布**：实时共享和协作
- **语音交互**：语音通话和语音控制
- **浏览器自动化**：自动执行 Web 操作

### 2.2 电子助理架构

```
┌─────────────────────────────────────────────────────────────────┐
│                    AI Assistant Gateway                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                   Channel Manager                         │    │
│  │  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ ┌───────┐ │    │
│  │  │ Web    │ │ Telegram│ │Discord │ │ Slack  │ │WhatsApp│ │    │
│  │  └────────┘ └────────┘ └────────┘ └────────┘ └───────┘ │    │
│  └─────────────────────────────────────────────────────────┘    │
│                              │                                   │
│                              ▼                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                   Multi-Agent Router                     │    │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────┐ │    │
│  │  │ Dev     │  │ Review  │  │ Project │  │Research│ │    │
│  │  │ Agent   │  │ Agent   │  │ Agent   │  │ Agent  │ │    │
│  │  └──────────┘  └──────────┘  └──────────┘  └────────┘ │    │
│  │                                                          │    │
│  │  每个 Agent 拥有独立的工作区、会话和工具权限              │    │
│  └─────────────────────────────────────────────────────────┘    │
│                              │                                   │
│                              ▼                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                    Skill Engine                          │    │
│  │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐    │    │
│  │  │ Code    │ │ Browser │ │ Voice   │ │ Canvas  │    │    │
│  │  │ Skill   │ │ Skill   │ │ Skill   │ │ Skill   │    │    │
│  │  └─────────┘ └─────────┘ └─────────┘ └─────────┘    │    │
│  │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐    │    │
│  │  │ Git     │ │ Terminal│ │ File    │ │ Search  │    │    │
│  │  │ Skill   │ │ Skill   │ │ Skill   │ │ Skill   │    │    │
│  │  └─────────┘ └─────────┘ └─────────┘ └─────────┘    │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 2.3 多平台消息集成

#### 2.3.1 支持的平台

| 平台 | 状态 | 说明 |
|------|------|------|
| Web Chat | 核心 | 内置 Web UI 聊天界面 |
| Telegram | 可选 | Bot 模式接入 |
| Discord | 可选 | Bot 模式接入 |
| Slack | 可选 | App 模式接入 |
| WhatsApp | 可选 | WhatsApp Business API |
| 钉钉 | 可选 | 钉钉机器人 |
| 飞书 | 可选 | 飞书机器人 |
| 企业微信 | 可选 | 企业微信应用 |

#### 2.3.2 跨平台消息配置

```json
{
  "channels": {
    "web": { "enabled": true },
    "telegram": {
      "enabled": true,
      "botToken": "${TELEGRAM_BOT_TOKEN}",
      "allowFrom": ["123456789"]
    },
    "discord": {
      "enabled": true,
      "token": "${DISCORD_BOT_TOKEN}",
      "dm": { "allowFrom": ["123456789012345678"] }
    }
  },
  "crossPlatform": {
    "allowAcrossProviders": true,
    "marker": { "enabled": true, "prefix": "[from {channel}] " }
  }
}
```

### 2.4 多 Agent 路由

#### 2.4.1 内置 Agent 类型

| Agent ID | 类型 | 工作区 | 专长 |
|----------|------|--------|------|
| dev | 开发代理 | workspace-dev | 代码开发、调试、重构 |
| review | 审查代理 | workspace-review | 代码审查、Bug 检测 |
| project | 项目代理 | workspace-project | 需求分析、任务管理 |
| research | 研究代理 | workspace-research | 技术调研、文档生成 |
| design | 设计代理 | workspace-design | UI 设计、架构设计 |

#### 2.4.2 Agent 路由配置

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.ai-assistant/workspace",
      "repoRoot": "~/Projects",
      "skipBootstrap": false,
      "bootstrapMaxChars": 20000
    },
    "list": [
      { "id": "dev", "default": true, "workspace": "~/.ai-assistant/workspace-dev" },
      { "id": "review", "workspace": "~/.ai-assistant/workspace-review" },
      { "id": "project", "workspace": "~/.ai-assistant/workspace-project" },
      { "id": "research", "workspace": "~/.ai-assistant/workspace-research" }
    ]
  },
  "bindings": [
    { "agentId": "dev", "match": { "channel": "web", "intent": "code" } },
    { "agentId": "review", "match": { "channel": "webhook", "type": "pull_request" } },
    { "agentId": "project", "match": { "channel": "web", "intent": "requirement" } }
  ]
}
```

#### 2.4.3 Agent 隔离机制

```json
{
  "agents": {
    "list": [{
      "id": "public",
      "workspace": "~/.ai-assistant/workspace-public",
      "sandbox": {
        "mode": "all",
        "scope": "agent",
        "workspaceAccess": "none"
      },
      "tools": {
        "allow": ["sessions_list", "sessions_history", "sessions_send"],
        "deny": ["read", "write", "exec", "browser"]
      }
    }]
  }
}
```

---

## 3. 技能系统 (Skill Engine)

### 3.1 技能架构

```
┌─────────────────────────────────────────────────────────────────┐
│                      Skill Engine                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│   ┌──────────────────────────────────────────────────────────┐  │
│   │                    Skill Registry                         │  │
│   │  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ ┌───────┐ │  │
│   │  │ Browser│ │ Canvas │ │ Voice  │ │ Git    │ │Terminal│ │  │
│   │  │Skill   │ │ Skill  │ │ Skill  │ │ Skill  │ │ Skill │ │  │
│   │  └────────┘ └────────┘ └────────┘ └────────┘ └───────┘ │  │
│   │  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ ┌───────┐ │  │
│   │  │ File   │ │ Search │ │ Code   │ │ Deploy │ │Market │ │  │
│   │  │ Skill  │ │ Skill  │ │ Skill  │ │ Skill  │ │ Skill │ │  │
│   │  └────────┘ └────────┘ └────────┘ └────────┘ └───────┘ │  │
│   └──────────────────────────────────────────────────────────┘  │
│                              │                                   │
│                              ▼                                   │
│   ┌──────────────────────────────────────────────────────────┐  │
│   │                    Skill Loader                           │  │
│   │   动态加载技能模块，支持第三方扩展                          │  │
│   └──────────────────────────────────────────────────────────┘  │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 3.2 内置技能

#### 3.2.1 Browser Skill（浏览器自动化）

```json
{
  "browser": {
    "executablePath": "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
    "headless": false,
    "userDataDir": "~/.ai-assistant/browser-data"
  }
}
```

**Browser Skill 能力**：
- 网页浏览和截图
- 表单自动填写
- 网页数据抓取
- UI 自动化测试

#### 3.2.2 Canvas Skill（实时画布）

```json
{
  "canvasHost": {
    "enabled": true,
    "port": 18793,
    "root": "/var/ai-assistant/canvas",
    "liveReload": true
  }
}
```

**Canvas Skill 能力**：
- 实时 HTML 预览
- 多人协作编辑
- 屏幕共享
- 实时演示

#### 3.2.3 Voice Skill（语音交互）

```json
{
  "voice": {
    "provider": "elevenlabs",
    "apiKey": "${ELEVENLABS_API_KEY}",
    "tts": {
      "model": "eleven_multilingual_v2",
      "voice": "default"
    },
    "stt": {
      "provider": "openai",
      "model": "whisper-1"
    },
    "wakeWord": "Hey Assistant"
  }
}
```

**Voice Skill 能力**：
- 语音唤醒（Wake Word）
- 语音转文字（STT）
- 文字转语音（TTS）
- 连续语音对话（Talk Mode）
- VoIP 语音通话

#### 3.2.4 Git Skill

**Git Skill 能力**：
- 仓库克隆和同步
- 分支管理
- 提交和推送
- PR/MR 管理
- 代码对比

#### 3.2.5 Terminal Skill

**Terminal Skill 能力**：
- 命令执行
- Shell 会话管理
- 脚本自动化
- 进程管理

#### 3.2.6 File Skill

**File Skill 能力**：
- 文件读写
- 目录操作
- 文件搜索
- 压缩解压

### 3.3 技能市场 (Skill Market)

```
┌─────────────────────────────────────────────────────────────────┐
│                      Skill Market                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│   官方技能                    │   第三方技能                       │
│   ┌─────────────────────┐    │   ┌─────────────────────┐        │
│   │ • Code Helper       │    │   │ • Jira Integration  │        │
│   │ • Browser Auto      │    │   │ • Notion Sync       │        │
│   │ • Voice Assistant   │    │   │ • Slack Workflow    │        │
│   │ • Canvas Editor     │    │   │ • GitHub Actions    │        │
│   │ • Database Tools    │    │   │ • AWS Management    │        │
│   └─────────────────────┘    │   └─────────────────────┘        │
│                                                                  │
│   安装命令: /skill install <skill-name>                          │
│   配置: ~/.ai-assistant/skills/<skill>/config.json              │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 4. 会话管理 (Session Management)

### 4.1 会话存储

会话数据使用 JSONL 格式存储：

```
~/.ai-assistant/sessions/
├── session-uuid-1/
│   ├── metadata.json      # 会话元数据
│   ├── messages.jsonl     # 消息历史（树形结构）
│   └── context.json       # 上下文缓存
├── session-uuid-2/
│   └── ...
```

### 4.2 消息树结构

```json
{
  "id": "msg-uuid",
  "parent_id": "msg-parent-uuid",
  "role": "user|assistant|system",
  "content": "消息内容",
  "timestamp": "2024-01-01T00:00:00Z",
  "metadata": {
    "model": "gpt-4o",
    "tokens": 1500,
    "attachments": []
  }
}
```

### 4.3 上下文管理

- **上下文预热**：提前加载会话文件
- **历史限制**：自动截断超长对话
- **上下文压缩**：智能总结旧对话

---

## 5. 实时画布 (Canvas)

### 5.1 画布架构

```
┌─────────────────────────────────────────────────────────────────┐
│                      Canvas Service                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│   ┌─────────────┐    WebSocket    ┌─────────────┐              │
│   │  Presenter │ ◄──────────────► │  Viewer(s)  │              │
│   │  (Agent)    │                 │  (Users)    │              │
│   └──────┬──────┘                 └─────────────┘              │
│          │                                                        │
│          │ HTTP / WebSocket                                       │
│          ▼                                                        │
│   ┌─────────────────────────────────────────────────────────┐    │
│   │              Canvas Server (Port: 18793)                │    │
│   │  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │    │
│   │  │  HTML       │  │  Static     │  │  Live Reload    │   │    │
│   │  │  Renderer   │  │  File       │  │  Hub            │   │    │
│   │  └─────────────┘  └─────────────┘  └─────────────────┘   │    │
│   └─────────────────────────────────────────────────────────┘    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 5.2 画布功能

| 功能 | 说明 |
|------|------|
| HTML 预览 | 实时渲染 HTML/CSS/JS |
| 屏幕共享 | 实时共享屏幕给用户 |
| 远程控制 | 用户可以远程协助 |
| 标注工具 | 屏幕标注和绘图 |
| 录制回放 | 记录操作过程 |

---

## 6. 语音服务 (Voice Service)

### 6.1 语音架构

```
┌─────────────────────────────────────────────────────────────────┐
│                      Voice Service                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│   ┌─────────────────────────────────────────────────────────┐    │
│   │                   Voice Gateway                          │    │
│   │  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │    │
│   │  │  Wake Word  │  │    STT      │  │      TTS        │   │    │
│   │  │  Detector   │  │  (Whisper)  │  │  (ElevenLabs)   │   │    │
│   │  └─────────────┘  └─────────────┘  └─────────────────┘   │    │
│   └─────────────────────────────────────────────────────────┘    │
│                              │                                   │
│                              ▼                                   │
│   ┌─────────────────────────────────────────────────────────┐    │
│   │                   Voice Channels                        │    │
│   │  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │    │
│   │  │   macOS     │  │   iOS       │  │    Android      │   │    │
│   │  │  (AVFAudio) │  │ (AVFoundation)│  │ (AudioTrack)   │   │    │
│   │  └─────────────┘  └─────────────┘  └─────────────────┘   │    │
│   └─────────────────────────────────────────────────────────┘    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 6.2 语音能力

| 能力 | 平台 | 说明 |
|------|------|------|
| Voice Wake | macOS/iOS | 语音唤醒词 |
| Talk Mode | Android | 连续语音对话 |
| VoIP Call | 全平台 | 语音通话 |
| TTS | 全平台 | 文字转语音 |

---

## 7. 前端架构 (Frontend)

### 7.1 技术选型
- **框架**：React 19
- **语言**：TypeScript 5.x
- **构建工具**：Vite 5.x
- **样式**：Tailwind CSS 3.x
- **状态管理**：Zustand / React Query
- **路由**：React Router 6.x
- **UI 组件**：shadcn/ui + Radix UI

### 7.2 前端模块结构

```
frontend/
├── public/
├── src/
│   ├── components/
│   │   ├── ui/              # UI 基础组件
│   │   ├── layout/          # 布局组件
│   │   ├── ide/             # IDE 相关组件
│   │   └── assistant/       # 电子助理组件
│   │       ├── Chat.tsx     # 聊天界面
│   │       ├── Canvas.tsx   # 实时画布
│   │       ├── Voice.tsx    # 语音控制
│   │       └── Skills.tsx   # 技能面板
│   ├── features/
│   │   ├── tasks/           # 智能任务
│   │   ├── projects/        # 项目管理
│   │   ├── ide/             # 在线开发环境
│   │   ├── review/          # 代码审查
│   │   ├── team/            # 团队协作
│   │   └── assistant/       # 电子助理
│   │       ├── channels/    # 渠道管理
│   │       ├── agents/      # Agent 配置
│   │       ├── skills/      # 技能管理
│   │       └── sessions/    # 会话管理
│   ├── hooks/
│   ├── services/
│   │   ├── api.ts          # API 服务
│   │   ├── ws.ts           # WebSocket 服务
│   │   └── assistant.ts    # 电子助理服务
│   ├── stores/
│   │   ├── assistant.ts    # 电子助理状态
│   │   ├── session.ts      # 会话状态
│   │   └── skill.ts        # 技能状态
│   ├── types/
│   └── utils/
├── package.json
└── vite.config.ts
```

---

## 8. 后端架构 (Backend)

### 8.1 技术选型
- **语言**：Go 1.21+
- **框架**：Echo v4
- **ORM**：Ent
- **数据库**：PostgreSQL 16
- **缓存**：Redis 7
- **文件存储**：S3 / MinIO

### 8.2 后端模块结构

```
backend/
├── cmd/
│   ├── server/              # 主服务入口
│   └── gateway/            # 电子助理网关入口
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── router/
│   ├── service/
│   │   ├── task/
│   │   ├── project/
│   │   ├── ai/
│   │   ├── ide/
│   │   ├── review/
│   │   ├── team/
│   │   └── assistant/      # 电子助理服务
│   │       ├── channel/    # 渠道管理
│   │       ├── agent/      # Agent 管理
│   │       ├── skill/      # 技能引擎
│   │       ├── session/    # 会话管理
│   │       ├── canvas/     # 画布服务
│   │       └── voice/      # 语音服务
│   ├── repository/
│   │   ├── ent/
│   │   └── cache/
│   ├── model/
│   └── pkg/
│       ├── ai/             # AI 模型客户端
│       ├── git/            # Git 操作
│       ├── docker/         # Docker 操作
│       ├── browser/        # 浏览器控制
│       └── voice/          # 语音处理
├── pkg/
│   ├── channel/            # 消息渠道
│   │   ├── telegram/
│   │   ├── discord/
│   │   ├── slack/
│   │   └── whatsapp/
│   └── skills/             # 技能实现
│       ├── browser/
│       ├── canvas/
│       ├── voice/
│       └── terminal/
├── migrations/
└── go.mod
```

---

## 9. AI 模型集成

### 9.1 AI 网关架构

```
                    ┌─────────────────┐
                    │   AI Gateway    │
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

### 9.2 支持的 AI 模型

| 模型 | Provider | 说明 |
|------|----------|------|
| GPT-4o | OpenAI | 主力模型 |
| GPT-4o-mini | OpenAI | 轻量级模型 |
| Claude 3.5 Sonnet | Anthropic | 代码能力强 |
| Claude 3.5 Haiku | Anthropic | 轻量级模型 |
| Qwen-Max | 通义千问 | 阿里大模型 |
| GLM-4-Plus | 智谱 | 智谱大模型 |
| SiliconFlow | 硅基流动 | 聚合 API |

---

## 10. 数据库设计

### 10.1 核心实体

**User（用户）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| email | string | 邮箱 |
| name | string | 名称 |
| avatar | string | 头像 URL |
| created_at | timestamp | 创建时间 |

**Team（团队）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| name | string | 团队名称 |
| owner_id | uuid | 所有者 ID |

**Agent（电子助理实例）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| name | string | Agent 名称 |
| type | enum | Agent 类型 |
| config | jsonb | Agent 配置 |
| workspace | string | 工作区路径 |
| team_id | uuid | 团队 ID |

**Channel（消息渠道）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| type | enum | 渠道类型 |
| config | jsonb | 渠道配置 |
| enabled | bool | 是否启用 |
| team_id | uuid | 团队 ID |

**Skill（技能）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| name | string | 技能名称 |
| version | string | 版本号 |
| config | jsonb | 技能配置 |
| enabled | bool | 是否启用 |

**Session（会话）**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uuid | 主键 |
| agent_id | uuid | Agent ID |
| channel | string | 来源渠道 |
| account_id | string | 账户标识 |
| status | enum | 会话状态 |
| created_at | timestamp | 创建时间 |

---

## 11. 安全架构

### 11.1 认证授权
- **JWT Token**：无状态认证
- **Refresh Token**：Token 续期
- **OAuth 2.0**：第三方登录
- **RBAC**：基于角色的权限控制

### 11.2 工作区隔离
```json
{
  "tools": {
    "fs": {
      "workspaceOnly": true
    }
  }
}
```

### 11.3 沙箱模式
- **off**：关闭沙箱
- **all**：所有工具受限
- **agent**：按 Agent 独立隔离

---

## 12. 部署架构

### 12.1 Docker Compose 部署

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
      - DATABASE_URL=postgresql://user:pass@db:5432/ai_coding
      - REDIS_URL=redis://redis:6379

  gateway:
    build: ./backend
    command: ./gateway
    ports:
      - "8081:8081"
    environment:
      - TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}
      - DISCORD_BOT_TOKEN=${DISCORD_BOT_TOKEN}

  canvas:
    build: ./canvas-service
    ports:
      - "18793:18793"

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

### 12.2 Kubernetes 部署

支持 Helm Chart 部署：
- Horizontal Pod Autoscaler
- Rolling Update
- ConfigMap / Secret 管理
- Service Mesh (Istio)
