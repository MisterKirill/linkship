import { BrowserRouter, Link, Routes } from 'react-router-dom'
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <div className="wrapper">
        <nav>
          <Link to="/" className="logo">Linkship</Link>

          <Link to="/login" className="nav-link">Log In</Link>
        </nav>

        <main>
          <Routes>

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
