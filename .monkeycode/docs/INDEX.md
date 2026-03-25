# AI 研发助手 - 项目概述

## 1. 项目简介

### 项目名称
**AI 研发助手** (AI Coding Assistant)

### 项目定位
企业级 AI 开发平台，融合 OpenClaw 设计理念，覆盖需求 → 设计 → 开发 → 代码审查全流程。

### 核心目标
通过自然语言交互，让 AI 完成从需求分析、技术设计、代码开发到代码审查的完整开发流程。

### 目标用户
- 研发团队（开发人员、测试人员）
- 技术管理者（项目经理、架构师）
- 需要提升研发效率的企业团队

---

## 2. 核心功能

### 2.1 电子助理（核心创新）

融合 OpenClaw 的多平台消息集成能力，提供统一的 AI 交互入口。

**功能特性**：
- **多平台消息集成**：Web、Telegram、Discord、Slack、WhatsApp 等
- **多 Agent 路由**：智能分发请求到专用 Agent（开发/审查/研究）
- **技能系统**：可扩展的能力体系（浏览器自动化、文件操作、Git 等）
- **实时画布**：实时共享和协作编辑
- **语音交互**：语音控制、语音通话、TTS/STT
- **会话管理**：多会话上下文、跨会话记忆

### 2.2 智能任务
用自然语言描述需求，AI 自动完成开发、设计或代码审查。

**功能特性**：
- 自然语言任务描述
- 支持多种 AI 模型
- Git 仓库关联与代码导入
- 代码生成与优化

### 2.3 项目管理
关联 Git 仓库，管理项目需求和任务。

**功能特性**：
- 需求文档管理
- 设计任务创建与跟踪
- 开发任务分配
- 进度追踪与统计

### 2.4 在线开发环境
提供完整的在线开发环境。

**功能特性**：
- **在线 IDE**：支持多语言语法高亮的代码编辑器
- **Web 终端**：支持多会话的浏览器终端
- **文件管理**：在线浏览、编辑、上传、下载文件
- **在线预览**：一键预览 Web 服务运行效果

### 2.5 代码审查
配置 Git 机器人，自动审查 PR/MR。

**功能特性**：
- GitHub/GitLab/Gitee Webhook 集成
- 智能代码分析
- 审查意见生成

### 2.6 团队协作
团队管理员可管理成员、分配资源。

**功能特性**：
- 团队成员管理
- 资源分配（宿主机、镜像、AI 模型）
- 权限控制

---

## 3. 技术架构

### 3.1 技术栈

| 层级 | 技术 |
|------|------|
| 前端 | React 19 + TypeScript + Vite + Tailwind CSS |
| 后端 | Go + Echo + Ent ORM |
| 数据库 | PostgreSQL 16 + Redis 7 |
| 桌面端 | Electron 35 |
| 移动端 | Capacitor 7（Android / iOS） |

### 3.2 核心模块

| 模块 | 说明 |
|------|------|
| Gateway | 电子助理网关，统一入口 |
| Channel Manager | 多平台消息渠道管理 |
| Multi-Agent Router | 多 Agent 智能路由 |
| Skill Engine | 技能引擎，支持扩展 |
| Session Manager | 会话管理与上下文 |
| Canvas Service | 实时画布服务 |
| Voice Service | 语音服务 |

### 3.3 项目结构

```
ai-coding-assistant/
├── frontend/          # Web 前端（React + TypeScript）
├── backend/           # 后端服务（Go）
│   ├── cmd/server/   # API 服务
│   └── cmd/gateway/  # 电子助理网关
├── desktop/           # 桌面客户端（Electron）
└── mobile/           # 移动客户端（Capacitor）
```

---

## 4. 电子助理详解

### 4.1 多平台消息

| 平台 | 状态 | 说明 |
|------|------|------|
| Web Chat | 核心 | 内置 Web UI |
| Telegram | 可选 | Bot 模式 |
| Discord | 可选 | Bot 模式 |
| Slack | 可选 | App 模式 |
| WhatsApp | 可选 | Business API |
| 钉钉 | 可选 | 机器人 |
| 飞书 | 可选 | 机器人 |

### 4.2 多 Agent 系统

| Agent | 专长 | 工作区 |
|--------|------|--------|
| dev | 代码开发 | workspace-dev |
| review | 代码审查 | workspace-review |
| project | 需求分析 | workspace-project |
| research | 技术调研 | workspace-research |

### 4.3 技能市场

| 类别 | 技能 |
|------|------|
| 开发 | Code Helper, Git, Terminal, File |
| 自动化 | Browser, Deploy |
| 协作 | Canvas, Voice |
| 集成 | Jira, Notion, Slack, GitHub Actions |

---

## 5. 里程碑规划

### Phase 1: 基础框架
- [ ] 项目脚手架
- [ ] 前端基础架构
- [ ] 后端基础架构
- [ ] 数据库设计

### Phase 2: 电子助理核心
- [ ] Gateway 网关
- [ ] Web Chat 界面
- [ ] 多 Agent 路由
- [ ] 会话管理

### Phase 3: 技能系统
- [ ] Skill Engine
- [ ] Browser Skill
- [ ] Git Skill
- [ ] Terminal Skill

### Phase 4: 高级功能
- [ ] Canvas 实时画布
- [ ] Voice 语音服务
- [ ] 多平台集成

### Phase 5: 客户端扩展
- [ ] 桌面客户端
- [ ] 移动客户端

---

## 6. 文档结构

| 文档 | 说明 |
|------|------|
| [INDEX.md](./INDEX.md) | 项目概述 |
| [ARCHITECTURE.md](./ARCHITECTURE.md) | 系统架构设计 |
| [INTERFACES.md](./INTERFACES.md) | 接口定义 |
| [DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md) | 开发者指南 |
