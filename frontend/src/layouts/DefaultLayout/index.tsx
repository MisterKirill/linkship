import { Link, NavLink, Outlet } from 'react-router-dom'
import classes from './style.module.css'

function DefaultLayout() {
  return (
    <div className={classes.wrapper}>
      <nav className={classes.nav}>
        <Link to="/" className={classes.logo}>
          Linkship
        </Link>

        <div className={classes.navLinks}>
          <NavLink to="/register" className={({ isActive }) => isActive ? classes.navLinkActive : classes.navLink}>
            Создать аккаунт
          </NavLink>
          <NavLink to="/login" className={({ isActive }) => isActive ? classes.navLinkActive : classes.navLink}>
            Войти
          </NavLink>
        </div>
      </nav>

      <main className={classes.main}>
        <Outlet />
      </main>

      <footer className={classes.footer}>
        <span>
          Made with ❤️ by <a href="https://github.com/MisterKirill" className="link">Mister Kirill</a>.<br />
          Source code can be found on <a href="https://github.com/MisterKirill/linkship" className="link">GitHub</a>.
        </span>
      </footer>
    </div>
  )
}

export default DefaultLayout
