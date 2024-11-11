import { Outlet } from 'react-router-dom'
import './style.css'

function ProfileLayout() {
  return (
    <main className="profile-content">
      <Outlet />
    </main>
  )
}

export default ProfileLayout
