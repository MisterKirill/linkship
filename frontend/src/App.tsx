import { BrowserRouter, Link, Routes } from 'react-router-dom'
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link to="/" className="logo">Linkship</Link>

        <Link to="/login" className="nav-link">Log In</Link>
      </nav>
      <main>
        <Routes>

        </Routes>
      </main>
    </BrowserRouter>
  )
}

export default App
