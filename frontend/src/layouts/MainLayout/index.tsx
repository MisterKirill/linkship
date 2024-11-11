import { Link, NavLink, Outlet } from 'react-router-dom'
import './style.css'

function MainLayout() {
  return (
    <div className="wrapper">
      <nav>
        <Link to="/" className="logo">Linkship</Link>

        <NavLink to="/login" className={({ isActive }) => isActive ? "nav-link nav-link-active" : "nav-link"}>
          Log In
        </NavLink>
      </nav>

      <main className="main-content">
        <Outlet />
      </main>

      <footer>
        Made with ❤️ by <a href="https://github.com/MisterKirill" className="text-link">Mister Kirill</a>.<br />
        Source code can be found on <a href="https://github.com/MisterKirill/linkship" className="text-link">GitHub</a>.
      </footer>
    </div>
  )
}

export default MainLayout
