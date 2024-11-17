import { useEffect, useState } from 'react'
import { Link, NavLink, Outlet } from 'react-router-dom'
import './style.css'

function MainLayout() {
  const [username, setUsername] = useState<string | null>(null)

  useEffect(() => {
    const storageUsername = localStorage.getItem('username')

    if (storageUsername) {
      setUsername(storageUsername)
    }
  }, [])

  return (
    <div className="wrapper">
      <nav className="nav">
        <Link to="/" className="logo">Linkship</Link>

        <div className="nav-links">
          {username ? (
            <NavLink to={`/${username}`} className="nav-link">
              {username}
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
            </NavLink>
          ) : (
            <>
              <NavLink to="/register" className={({ isActive }) => isActive ? "nav-link nav-link-active" : "nav-link"}>
                Register
              </NavLink>
              <NavLink to="/login" className={({ isActive }) => isActive ? "nav-link nav-link-active" : "nav-link"}>
                Log In
              </NavLink>
            </>
          )}
        </div>
      </nav>

      <main className="main-content">
        <Outlet />
      </main>

      <footer className="footer">
        Made with ❤️ by <a href="https://github.com/MisterKirill" className="text-link">Mister Kirill</a>.<br />
        Source code can be found on <a href="https://github.com/MisterKirill/linkship" className="text-link">GitHub</a>.
      </footer>
    </div>
  )
}

export default MainLayout
