import React from 'react';
import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'; // ✅ Fix: Devtools import
import { ToastContainer } from 'react-toastify';


import MainLayout from './pages/MainLayout';
import About from './pages/About';
import Home from './pages/Home';
import Farms from './pages/farms/Farms';
import PostFarm from './pages/farms/PostFarm';
import SignIn from './pages/Authentication/SignIn'; // ✅ Fix: Import SignIn component
import SignUp from './pages/Authentication/SignUp';




const router = createBrowserRouter([
  {
    path: '/',
    element: <MainLayout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: 'about',
        element: <About />,
      },
      {
        path: 'farms',
        element: <Farms />,
      },
      {
        path: 'farms/post',
        element: <PostFarm />,
      },
      {
        path: 'auth/signin',
        element: <SignIn />,
      },
      {
        path: 'auth/signup',
        element: <SignUp />,
      },
    ],
  },
]);

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
       <ToastContainer position="top-right" />
      <ReactQueryDevtools initialIsOpen={false} /> {/* ✅ Corrected DevTools usage */}
    </QueryClientProvider>
  );
}

export default App;
