import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router";
import './index.css'
import App from './App.tsx'
import SearchPage from './SearchPage.tsx'
import ReservationPage from './ReservationPage.tsx'
import HotelPage from './HotelPage.tsx'
import HotelsPage from './HotelsPage.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />}>
          <Route path="search" element={<SearchPage />} />
          <Route path="reservation" element={<ReservationPage />} />
          <Route path="hotels" element={<HotelsPage />} />
          <Route path="hotels/:hotelId" element={<HotelPage />} />
        </Route>
      </Routes>
    </BrowserRouter >
  </StrictMode>,
)
