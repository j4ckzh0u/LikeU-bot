# AI 研发助手 - 接口定义

## 1. API 概述

### 1.1 基本信息
- **基础 URL**：`/api/v1`
- **认证方式**：JWT Bearer Token
- **请求格式**：`application/json`
- **响应格式**：`application/json`

### 1.2 通用响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| code | int | 状态码，0 表示成功 |
| message | string | 状态信息 |
| data | object | 响应数据 |

### 1.3 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

---

## 2. 认证接口

### 2.1 用户登录
```
POST /api/v1/auth/login
```

**请求体**：
```json
{
  "email": "user@example.com",
  "password": "password"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_in": 86400,
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "name": "张三"
    }
  }
}
```

### 2.2 OAuth 登录
```
GET /api/v1/auth/oauth/:provider
```

**Provider**：github, gitlab, gitee

**响应**：重定向到第三方授权页面

### 2.3 OAuth 回调
```
GET /api/v1/auth/oauth/:provider/callback
```

**Query 参数**：
- `code`：授权码
- `state`：CSRF 状态

### 2.4 刷新 Token
```
POST /api/v1/auth/refresh
```

**请求体**：
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### 2.5 退出登录
```
POST /api/v1/auth/logout
```

---

## 3. 用户接口

### 3.1 获取当前用户
```
GET /api/v1/user/me
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "张三",
    "avatar": "https://...",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 3.2 更新用户信息
```
PUT /api/v1/user/me
```

**请求体**：
```json
{
  "name": "张三",
  "avatar": "https://..."
}
```

---

## 4. 团队接口

### 4.1 创建团队
```
POST /api/v1/teams
```

**请求体**：
```json
{
  "name": "研发团队"
}
```

### 4.2 团队列表
```
GET /api/v1/teams
```

**Query 参数**：
- `page`：页码
- `page_size`：每页数量

### 4.3 获取团队详情
```
GET /api/v1/teams/:id
```

### 4.4 更新团队
```
PUT /api/v1/teams/:id
```

### 4.5 删除团队
```
DELETE /api/v1/teams/:id
```

### 4.6 团队成员管理

**添加成员**：
```
POST /api/v1/teams/:id/members
```

**请求体**：
```json
{
  "user_id": "uuid",
  "role": "member"
}
```

**角色**：owner, admin, member

**移除成员**：
```
DELETE /api/v1/teams/:id/members/:user_id
```

**成员列表**：
```
GET /api/v1/teams/:id/members
```

---

## 5. 项目接口

### 5.1 创建项目
```
POST /api/v1/projects
```

**请求体**：
```json
{
  "name": "我的项目",
  "description": "项目描述",
  "team_id": "uuid",
  "git_url": "https://github.com/xxx/yyy"
}
```

### 5.2 项目列表
```
GET /api/v1/projects
```

**Query 参数**：
- `team_id`：团队 ID
- `page`：页码
- `page_size`：每页数量

### 5.3 获取项目详情
```
GET /api/v1/projects/:id
```

### 5.4 更新项目
```
PUT /api/v1/projects/:id
```

### 5.5 删除项目
```
DELETE /api/v1/projects/:id
```

### 5.6 同步 Git 仓库
```
POST /api/v1/projects/:id/sync
```

---

## 6. 任务接口

### 6.1 创建任务
```
POST /api/v1/tasks
```

**请求体**：
```json
{
  "title": "实现用户登录功能",
  "description": "用户可以通过邮箱密码登录系统",
  "type": "dev",
  "project_id": "uuid"
}
```

**任务类型**：
- `dev`：开发任务
- `design`：设计任务
- `review`：审查任务

### 6.2 任务列表
```
GET /api/v1/tasks
```

**Query 参数**：
- `project_id`：项目 ID
- `type`：任务类型
- `status`：任务状态
- `page`：页码
- `page_size`：每页数量

### 6.3 获取任务详情
```
GET /api/v1/tasks/:id
```

### 6.4 更新任务
```
PUT /api/v1/tasks/:id
```

### 6.5 删除任务
```
DELETE /api/v1/tasks/:id
```

### 6.6 任务状态流转

**状态**：pending → in_progress → completed

**更新状态**：
```
PUT /api/v1/tasks/:id/status
```

**请求体**：
```json
{
  "status": "in_progress"
}
```

---

## 7. AI 对话接口

### 7.1 发送消息
```
POST /api/v1/tasks/:id/messages
```

**请求体**：
```json
{
  "content": "帮我实现用户登录功能",
  "model": "gpt-4o"
}
```

**响应**（SSE 流式）：
```
Content-Type: text/event-stream

data: {"role": "assistant", "content": "好的，我"}

data: {"role": "assistant", "content": "来帮你"}

data: {"role": "assistant", "content": "实现这个功能。"}
```

### 7.2 获取对话历史
```
GET /api/v1/tasks/:id/messages
```

**Query 参数**：
- `page`：页码
- `page_size`：每页数量

### 7.3 清空对话历史
```
DELETE /api/v1/tasks/:id/messages
```

---

## 8. AI 模型接口

### 8.1 获取可用模型
```
GET /api/v1/ai/models
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "gpt-4o",
      "name": "GPT-4o",
      "provider": "openai",
      "enabled": true
    },
    {
      "id": "claude-3-5-sonnet",
      "name": "Claude 3.5 Sonnet",
      "provider": "anthropic",
      "enabled": true
    }
  ]
}
```

### 8.2 配置 API Key
```
POST /api/v1/ai/keys
```

**请求体**：
```json
{
  "provider": "openai",
  "api_key": "sk-xxx"
}
```

### 8.3 删除 API Key
```
DELETE /api/v1/ai/keys/:provider
```

---

## 9. 在线 IDE 接口

### 9.1 创建开发环境
```
POST /api/v1/ide/environments
```

**请求体**：
```json
{
  "project_id": "uuid",
  "image": "ubuntu:22.04"
}
```

### 9.2 获取开发环境
```
GET /api/v1/ide/environments/:id
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "uuid",
    "status": "running",
    "host": "ide.example.com",
    "port": 8080,
    "token": "xxx"
  }
}
```

### 9.3 关闭开发环境
```
DELETE /api/v1/ide/environments/:id
```

### 9.4 WebSocket 连接
```
WS /api/v1/ide/terminal/:environment_id
```

### 9.5 文件操作

**列出文件**：
```
GET /api/v1/ide/files
```

**读取文件**：
```
GET /api/v1/ide/files/:path
```

**写入文件**：
```
PUT /api/v1/ide/files/:path
```

**删除文件**：
```
DELETE /api/v1/ide/files/:path
```

---

## 10. 代码审查接口

### 10.1 创建审查任务
```
POST /api/v1/reviews
```

**请求体**：
```json
{
  "project_id": "uuid",
  "provider": "github",
  "pull_request_id": "123"
}
```

### 10.2 审查任务列表
```
GET /api/v1/reviews
```

**Query 参数**：
- `project_id`：项目 ID
- `status`：审查状态
- `page`：页码

### 10.3 获取审查详情
```
GET /api/v1/reviews/:id
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "uuid",
    "status": "completed",
    "provider": "github",
    "pull_request_url": "https://github.com/xxx/yyy/pull/123",
    "summary": {
      "total_files": 5,
      "total_comments": 12,
      "risk_level": "medium"
    },
    "comments": [
      {
        "file": "src/login.ts",
        "line": 42,
        "content": "建议：添加输入验证",
        "severity": "warning"
      }
    ],
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 10.4 配置 Webhook
```
POST /api/v1/reviews/webhook
```

**请求体**：
```json
{
  "provider": "github",
  "events": ["pull_request", "push"]
}
```

---

## 11. 资源接口

### 11.1 宿主机管理

**创建宿主机**：
```
POST /api/v1/resources/hosts
```

**宿主机列表**：
```
GET /api/v1/resources/hosts
```

**删除宿主机**：
```
DELETE /api/v1/resources/hosts/:id
```

### 11.2 镜像管理

**镜像列表**：
```
GET /api/v1/resources/images
```

**构建镜像**：
```
POST /api/v1/resources/images/build
```

---

## 12. 电子助理接口 (Assistant)

### 12.1 Agent 管理

**创建 Agent**：
```
POST /api/v1/assistant/agents
```

**请求体**：
```json
{
  "name": "开发助手",
  "type": "dev",
  "config": {
    "model": "gpt-4o",
    "temperature": 0.7,
    "max_tokens": 4000
  },
  "workspace": "/path/to/workspace"
}
```

**Agent 类型**：dev, review, project, research, custom

**Agent 列表**：
```
GET /api/v1/assistant/agents
```

**获取 Agent**：
```
GET /api/v1/assistant/agents/:id
```

**更新 Agent**：
```
PUT /api/v1/assistant/agents/:id
```

**删除 Agent**：
```
DELETE /api/v1/assistant/agents/:id
```

### 12.2 渠道管理 (Channel Manager)

**创建渠道**：
```
POST /api/v1/assistant/channels
```

**请求体**：
```json
{
  "type": "telegram",
  "name": "我的 Telegram",
  "config": {
    "botToken": "xxx",
    "allowFrom": ["123456789"]
  },
  "enabled": true
}
```

**渠道类型**：web, telegram, discord, slack, whatsapp, dingtalk, feishu

**渠道列表**：
```
GET /api/v1/assistant/channels
```

**更新渠道配置**：
```
PUT /api/v1/assistant/channels/:id
```

**启用/禁用渠道**：
```
PATCH /api/v1/assistant/channels/:id/status
```

**请求体**：
```json
{
  "enabled": true
}
```

### 12.3 会话管理 (Session)

**创建会话**：
```
POST /api/v1/assistant/sessions
```

**请求体**：
```json
{
  "agent_id": "uuid",
  "channel": "telegram",
  "account_id": "123456789"
}
```

**会话列表**：
```
GET /api/v1/assistant/sessions
```

**Query 参数**：
- `agent_id`：Agent ID
- `channel`：渠道类型
- `status`：会话状态

**获取会话详情**：
```
GET /api/v1/assistant/sessions/:id
```

**获取会话历史**：
```
GET /api/v1/assistant/sessions/:id/messages
```

**Query 参数**：
- `page`：页码
- `page_size`：每页数量
- `parent_id`：父消息 ID（获取回复链）

**发送消息**：
```
POST /api/v1/assistant/sessions/:id/messages
```

**请求体**：
```json
{
  "content": "帮我实现用户登录功能",
  "parent_id": "uuid"
}
```

**响应**（SSE 流式）：
```
Content-Type: text/event-stream

data: {"role": "assistant", "content": "好的，我"}

data: {"role": "assistant", "content": "来帮你"}

data: {"role": "assistant", "content": "实现这个功能。"}
```

**清空会话历史**：
```
DELETE /api/v1/assistant/sessions/:id/messages
```

### 12.4 技能管理 (Skill Engine)

**技能列表**：
```
GET /api/v1/assistant/skills
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "browser",
        "name": "Browser",
        "version": "1.0.0",
        "description": "浏览器自动化技能",
        "enabled": true,
        "builtin": true
      },
      {
        "id": "git",
        "name": "Git",
        "version": "1.0.0",
        "description": "Git 操作技能",
        "enabled": true,
        "builtin": true
      }
    ],
    "total": 10
  }
}
```

**安装技能**：
```
POST /api/v1/assistant/skills/:id/install
```

**卸载技能**：
```
DELETE /api/v1/assistant/skills/:id/install
```

**技能配置**：
```
PUT /api/v1/assistant/skills/:id/config
```

**请求体**：
```json
{
  "browser": {
    "executablePath": "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
    "headless": false
  }
}
```

### 12.5 实时画布 (Canvas)

**创建画布**：
```
POST /api/v1/assistant/canvas
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "uuid",
    "url": "http://canvas.example.com/canvas/abc123",
    "wsUrl": "ws://canvas.example.com/canvas/abc123"
  }
}
```

**获取画布**：
```
GET /api/v1/assistant/canvas/:id
```

**共享画布**：
```
POST /api/v1/assistant/canvas/:id/share
```

**请求体**：
```json
{
  "user_ids": ["uuid1", "uuid2"],
  "readonly": false
}
```

**更新画布内容**：
```
PUT /api/v1/assistant/canvas/:id/content
```

**请求体**：
```json
{
  "html": "<html>...",
  "css": "...",
  "js": "..."
}
```

**WebSocket 画布连接**：
```
WS /api/v1/assistant/canvas/:id/ws
```

**消息格式**：
```json
{
  "type": "update|cursor|annotation",
  "data": {}
}
```

### 12.6 语音服务 (Voice)

**发起语音通话**：
```
POST /api/v1/assistant/voice/call
```

**请求体**：
```json
{
  "session_id": "uuid",
  "to": "+15555550123",
  "mode": "voice"
}
```

**语音通话状态**：
```
GET /api/v1/assistant/voice/calls/:id
```

**继续通话（发送消息）**：
```
POST /api/v1/assistant/voice/calls/:id/continue
```

**请求体**：
```json
{
  "message": "有什么问题吗？"
}
```

**结束通话**：
```
DELETE /api/v1/assistant/voice/calls/:id
```

**设置语音配置**：
```
PUT /api/v1/assistant/voice/config
```

**请求体**：
```json
{
  "provider": "elevenlabs",
  "apiKey": "xxx",
  "tts": {
    "model": "eleven_multilingual_v2",
    "voice": "default"
  },
  "stt": {
    "provider": "openai",
    "model": "whisper-1"
  },
  "wakeWord": {
    "enabled": true,
    "word": "Hey Assistant"
  }
}
```

### 12.7 跨平台消息

**配置跨平台消息**：
```
PUT /api/v1/assistant/cross-platform
```

**请求体**：
```json
{
  "allowAcrossProviders": true,
  "marker": {
    "enabled": true,
    "prefix": "[from {channel}] "
  }
}
```

---

## 13. 错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 缺少必填参数 |
| 2001 | 未登录 |
| 2002 | Token 无效 |
| 2003 | 权限不足 |
| 3001 | 资源不存在 |
| 3002 | 资源已存在 |
| 4001 | 内部服务器错误 |
| 4002 | AI 服务调用失败 |
| 4003 | Git 服务调用失败 |
