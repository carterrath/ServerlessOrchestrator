//import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Home } from '../pages/Home';
import { Login } from '../pages/Login';
import { Developer } from '../pages/Developer';
import DeveloperOptions from '../pages/DeveloperOptions';
import { Microservices } from '../pages/Microservices';
import { NavBar } from '../components/NavBar';

export function Router() {
  return (
    <BrowserRouter>
    <NavBar isDeveloper={true} />
      <Routes>
        <Route path="/*"  element={<Home />} />
        <Route path="/Login"  element={<Login />} />
        <Route path="/Developer" element={<Developer />} />
        <Route path="/DeveloperOptions" element={<DeveloperOptions />} />
        <Route path="/Microservices" element={<Microservices />} />
    
      </Routes>
    </BrowserRouter>
  );
}