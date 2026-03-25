# AI 研发助手 - 开发者指南

## 1. 开发环境准备

### 1.1 系统要求

- **操作系统**：Linux / macOS / Windows (WSL2)
- **内存**：至少 16GB
- **磁盘**：至少 50GB 可用空间
- **Docker**：Docker 20.10+ / Docker Desktop

### 1.2 必需安装的软件

| 软件 | 版本 | 说明 |
|------|------|------|
| Go | 1.21+ | 后端开发 |
| Node.js | 18+ | 前端开发 |
| pnpm | 8+ | 前端包管理 |
| Docker | 20.10+ | 本地运行 PostgreSQL、Redis |
| Git | 2.30+ | 版本控制 |

### 1.3 开发工具推荐

- **IDE**：VSCode / GoLand / WebStorm
- **终端**：iTerm2 / Windows Terminal
- **API 测试**：Postman / Insomnia

---

## 2. 项目克隆与初始化

### 2.1 克隆项目

```bash
git clone https://github.com/your-org/ai-coding-assistant.git
cd ai-coding-assistant
```

### 2.2 初始化 Git Submodules

```bash
git submodule update --init --recursive
```

### 2.3 启动依赖服务

使用 Docker Compose 启动 PostgreSQL 和 Redis：

```bash
docker-compose up -d db redis
```

---

## 3. 后端开发

### 3.1 项目结构

```
backend/
├── cmd/server/           # 主入口
├── internal/             # 内部包
│   ├── api/             # HTTP 层
│   ├── service/         # 业务逻辑
│   ├── repository/      # 数据访问
│   └── model/           # 数据模型
├── pkg/                 # 公共包
├── migrations/          # 数据库迁移
└── go.mod
```

### 3.2 环境变量配置

创建 `backend/.env` 文件：

```bash
# 数据库配置
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/ai_coding?sslmode=disable
DB_DRIVER=postgres

# Redis 配置
REDIS_URL=redis://localhost:6379/0

# JWT 配置
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRY=24h

# AI 模型配置
OPENAI_API_KEY=sk-your-key
ANTHROPIC_API_KEY=sk-ant-your-key

# Git 配置
GITHUB_WEBHOOK_SECRET=your-webhook-secret

# 文件存储
STORAGE_TYPE=local
STORAGE_PATH=./data/storage

# 开发模式
GIN_MODE=debug
```

### 3.3 数据库迁移

```bash
cd backend

# 安装 golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 运行迁移
make migrate-up

# 创建新迁移
make migrate-create name=add_user_table

# 回滚迁移
make migrate-down
```

### 3.4 启动后端服务

```bash
cd backend

# 安装依赖
go mod tidy

# 运行
go run ./cmd/server/main.go

# 或者使用 Makefile
make run
```

服务启动后，运行在 `http://localhost:8080`。

### 3.5 后端 API 文档

启动服务后访问：
- Swagger UI：`http://localhost:8080/swagger`
- API 端点：`http://localhost:8080/api/v1`

### 3.6 测试

```bash
# 运行所有测试
make test

# 运行指定包的测试
go test ./internal/service/...

# 生成测试覆盖率报告
make test-coverage
```

---

## 4. 前端开发

### 4.1 项目结构

```
frontend/
├── public/              # 静态资源
├── src/
│   ├── components/     # 公共组件
│   ├── features/       # 功能模块
│   ├── hooks/         # 自定义 Hooks
│   ├── services/      # API 服务
│   ├── stores/        # 状态管理
│   └── utils/        # 工具函数
├── package.json
└── vite.config.ts
```

### 4.2 环境变量配置

创建 `frontend/.env.local` 文件：

```bash
# API 地址
VITE_API_BASE_URL=http://localhost:8080/api/v1

# WebSocket 地址
VITE_WS_URL=ws://localhost:8080

# AI 服务地址
VITE_AI_GATEWAY_URL=http://localhost:8080
```

### 4.3 安装依赖

```bash
cd frontend

# 安装依赖
pnpm install

# 如果没有 pnpm，先安装
npm install -g pnpm
```

### 4.4 启动前端服务

```bash
cd frontend

# 开发模式
pnpm dev

# 构建生产版本
pnpm build

# 预览生产版本
pnpm preview
```

服务启动后，运行在 `http://localhost:3000`。

### 4.5 代码规范

```bash
# 运行 ESLint
pnpm lint

# 修复 ESLint 错误
pnpm lint:fix

# 运行 TypeScript 类型检查
pnpm typecheck
```

### 4.6 组件开发

使用 shadcn/ui 组件库：

```bash
# 添加新组件
pnpm add @radix-ui/react-dialog
npx shadcn-ui@latest add button

# 查看可用组件
npx shadcn-ui@latest show
```

---

## 5. 在线开发环境

### 5.1 开发环境架构

在线 IDE 使用以下组件：
- **Monaco Editor**：代码编辑
- **xterm.js**：终端模拟
- **code-server**：VSCode 后端

### 5.2 本地开发

开发环境需要 Docker 支持：

```bash
# 构建开发环境镜像
make ide-build

# 启动开发环境
make ide-up

# 停止开发环境
make ide-down
```

### 5.3 WebSocket 连接

IDE 使用 WebSocket 进行实时通信：

```javascript
const ws = new WebSocket('ws://localhost:8080/api/v1/ide/terminal/:env_id');
```

---

## 6. 代码审查功能

### 6.1 GitHub Webhook 配置

1. 在 GitHub 仓库设置中添加 Webhook
2. Payload URL：`https://your-domain.com/api/v1/reviews/webhook/github`
3. Content Type：`application/json`
4. Secret：与配置中的 `GITHUB_WEBHOOK_SECRET` 一致
5. Events：选择 `Pull requests`

### 6.2 本地测试 Webhook

使用 ngrok 进行本地测试：

```bash
ngrok http 8080
```

---

## 7. Docker Compose 开发环境

### 7.1 完整服务启动

```bash
# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f backend

# 停止所有服务
docker-compose down
```

### 7.2 服务端口

| 服务 | 端口 | 说明 |
|------|------|------|
| Frontend | 3000 | 前端 Web |
| Backend | 8080 | 后端 API |
| PostgreSQL | 5432 | 数据库 |
| Redis | 6379 | 缓存 |

---

## 8. 生产部署

### 8.1 构建镜像

```bash
# 构建前端
cd frontend && pnpm build

# 构建后端
cd backend && make build

# 构建 Docker 镜像
docker build -t ai-coding-assistant:latest .
```

### 8.2 Kubernetes 部署

使用 Helm Chart 部署：

```bash
# 添加 Helm 仓库
helm repo add my-repo https://charts.example.com

# 安装
helm install ai-coding my-repo/ai-coding-assistant

# 升级
helm upgrade ai-coding my-repo/ai-coding-assistant -f values.yaml
```

### 8.3 环境变量

生产环境需要设置以下环境变量：

```bash
# 数据库
DATABASE_URL=postgresql://user:pass@db:5432/prod?sslmode=require

# Redis
REDIS_URL=redis://redis:6379/0

# JWT
JWT_SECRET=<generate-strong-secret>

# AI API Keys
OPENAI_API_KEY=sk-xxx
ANTHROPIC_API_KEY=sk-ant-xxx

# 域名
DOMAIN=ai-coding.example.com
```

---

## 9. 开发规范

### 9.1 Git 提交规范

使用 Conventional Commits：

```
feat: 新功能
fix: 修复 bug
docs: 文档更新
style: 代码格式
refactor: 重构
test: 测试
chore: 构建/工具
```

示例：
```
feat(tasks): 添加任务评论功能
fix(auth): 修复 Token 刷新问题
docs(api): 更新用户接口文档
```

### 9.2 代码审查

- 所有代码需要通过 CI 检查
- 至少需要 1 人 review 通过才能合并
- 确保测试覆盖率不低于 70%

### 9.3 分支命名

```
YYYYMMDD-(feat|fix|chore|refactor)-short-description
```

示例：
```
260325-feat-add-user-login
260324-fix-task-status-update
```

---

## 10. 常见问题

### 10.1 数据库连接失败

确保 PostgreSQL 服务正在运行：
```bash
docker-compose up -d db
```

### 10.2 前端 API 请求失败

检查后端服务是否正常运行，并确认 `VITE_API_BASE_URL` 配置正确。

### 10.3 AI 模型调用失败

检查 API Key 配置是否正确，确认网络可以访问 AI 服务商。

### 10.4 Docker 构建失败

清理 Docker 缓存：
```bash
docker builder prune
```
