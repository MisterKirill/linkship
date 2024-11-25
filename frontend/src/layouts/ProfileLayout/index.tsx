import { Link, Outlet } from 'react-router-dom'
import './style.css'

function ProfileLayout() {
  return (
    <main className="profile-content">
      <Outlet />

      <Link to="/" className="powered-by">Powered by <b>Linkship</b></Link>
    </main>
  )
}

export default ProfileLayout
