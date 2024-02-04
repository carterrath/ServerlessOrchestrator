//import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Home } from '../pages/Home';
import { Login } from '../pages/Login';
import Developer from '../pages/Developer';
import DeveloperOptions from '../pages/DeveloperOptions';
export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/*"  element={<Home />} />
        <Route path="/Login"  element={<Login />} />
        <Route path="/Developer" element={<Developer />} />
        <Route path="DeveloperOptions" element={<DeveloperOptions />} />
        
      </Routes>
    </BrowserRouter>
  );
}