import { BrowserRouter, Route, Routes } from 'react-router-dom'
import DefaultLayout from './layouts/DefaultLayout'
import Home from './pages/Home'
import NotFound from './pages/NotFound'
import Login from './pages/Login'
import Register from './pages/Register'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<DefaultLayout />}>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="*" element={<NotFound />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
