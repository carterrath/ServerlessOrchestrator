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
export function Router() {
  return (
    <BrowserRouter>
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
          </Routes>
      </div>
    </BrowserRouter>
  );
}