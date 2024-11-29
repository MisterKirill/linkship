import { Link, Outlet } from 'react-router-dom'
import './style.css'

function DefaultLayout() {
  return (
    <div className="wrapper">
      <nav className="nav">
        <Link to="/" className="logo">
          Linkship
        </Link>

        <div className="links">
          <Link to="/register" className="nav-link">Register</Link>
          <Link to="/login" className="nav-link">Log In</Link>
        </div>
      </nav>

      <main className="main">
        <Outlet />
      </main>

      <footer className="footer">
        <span>
          Made with ❤️ by <a href="https://github.com/MisterKirill" className="link">Mister Kirill</a>.<br />
          Source code can be found on <a href="https://github.com/MisterKirill/linkship" className="link">GitHub</a>.
        </span>
      </footer>
    </div>
  )
}

export default DefaultLayout
