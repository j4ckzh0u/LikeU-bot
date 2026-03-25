import { Routes, Route, Navigate } from 'react-router-dom'
import { MainLayout } from '@/components/layout/MainLayout'
import { AssistantPage } from '@/features/assistant/pages/AssistantPage'
import { ProjectsPage } from '@/features/projects/pages/ProjectsPage'
import { TasksPage } from '@/features/tasks/pages/TasksPage'
import { ReviewPage } from '@/features/review/pages/ReviewPage'
import { TeamPage } from '@/features/team/pages/TeamPage'
import { LoginPage } from '@/features/auth/pages/LoginPage'
import { useAuthStore } from '@/stores/auth'

function PrivateRoute({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuthStore()
  return isAuthenticated ? children : <Navigate to="/login" replace />
}

function App() {
  return (
    <Routes>
      <Route path="/login" element={<LoginPage />} />
      <Route
        path="/"
        element={
          <PrivateRoute>
            <MainLayout />
          </PrivateRoute>
        }
      >
        <Route index element={<AssistantPage />} />
        <Route path="projects" element={<ProjectsPage />} />
        <Route path="tasks" element={<TasksPage />} />
        <Route path="review" element={<ReviewPage />} />
        <Route path="team" element={<TeamPage />} />
      </Route>
    </Routes>
  )
}

export default App
