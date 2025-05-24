import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import Login from './Login.jsx'
import Actividades from './Actividades.jsx'
import Layout from './Layout.jsx'
import Home from './Home.jsx'
import { BrowserRouter, Routes, Route } from "react-router";
import AddVinil from './AddVinil.jsx'
import Vinils from './Vinil.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
  <BrowserRouter>
    <Routes>

    <Route path="/" element={<Layout />}>
      <Route index element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/actividades" element={<Actividades />} />
      <Route path="/vinilos" element={<Vinils />} />
      <Route path="/addVinilos" element={<AddVinil />} />
    </Route>
    </Routes>
  </BrowserRouter>
  </StrictMode>,
)
