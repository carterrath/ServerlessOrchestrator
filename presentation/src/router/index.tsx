//import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Home } from '../pages/Home';
import { UploadMicroservice } from '../pages/UploadMicroservice';
import DeveloperOptions from '../pages/DeveloperOptions';
import { Microservices } from '../pages/Microservices';
import { NavBar } from '../components/NavBar';
import MainChoice from '../pages/GetStarted/index';
import DeveloperSignup from '../pages/GetStarted/DeveloperSignup';
import ConsumerSignup from '../pages/GetStarted/ConsumerSignup';
import DeveloperLogin from '../pages/GetStarted/DeveloperLogin';
import ConsumerLogin from '../pages/GetStarted/ConsumerLogin';
import ResetPassword from '../pages/GetStarted/ResetPassword';
import Recover from '../pages/GetStarted/Recover';
import { AuthProvider } from '../hooks/useAuth';
export function Router() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <NavBar isDeveloper={true} />
        <div className="min-h-[calc(100svh-64px)] w-full absolute top-22">
            <Routes>
              <Route path="/*"  element={<Home />} />
              <Route path="/UploadMicroservice" element={<UploadMicroservice />} />
              <Route path="/DeveloperOptions" element={<DeveloperOptions />} />
              <Route path="/Microservices" element={<Microservices />} />
              <Route path="/GetStarted" element={<MainChoice />} />
              <Route path="/developer-signup" element={<DeveloperSignup />} />
              <Route path="/consumer-signup" element={<ConsumerSignup />} />
              <Route path="/developer-login" element={<DeveloperLogin />} />
              <Route path="/consumer-login" element={<ConsumerLogin />} />
              <Route path="/recover-account" element={<Recover />} />
              <Route path="/reset-password" element={<ResetPassword />} />
            </Routes>
        </div>
      </AuthProvider>
    </BrowserRouter>
  );
}