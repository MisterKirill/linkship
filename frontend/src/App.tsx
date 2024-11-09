import './App.css'
import { BrowserRouter, Link, NavLink, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import LogIn from './pages/LogIn'
import NotFound from './pages/NotFound'

function App() {
  return (
    <BrowserRouter>
      <div className="wrapper">
        <nav>
          <Link to="/" className="logo">Linkship</Link>

          <NavLink to="/login" className={({ isActive }) => isActive ? "nav-link nav-link-active" : "nav-link"}>
            Log In
          </NavLink>
        </nav>

        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<LogIn />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </main>

        <footer>
          Made with ❤️ by <a href="https://github.com/MisterKirill" className="text-link">Mister Kirill</a>.<br />
          Source code can be found on <a href="https://github.com/MisterKirill/linkship" className="text-link">GitHub</a>.
        </footer>
      </div>
    </BrowserRouter>
  )
}

export default App
