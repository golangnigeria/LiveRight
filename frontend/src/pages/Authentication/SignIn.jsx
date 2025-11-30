import React from 'react';
import { Link } from 'react-router-dom';
import { Mail, Lock } from 'lucide-react';
import logo from '../../assets/nnefarm.png';

export default function SignIn() {
  return (
    <div className="min-h-screen pt-24 bg-gradient-to-br from-[#FAF6EF] to-[#FFFFFF] px-4">
      <div className="max-w-4xl mx-auto bg-white rounded-2xl shadow-lg overflow-hidden grid md:grid-cols-2">
        
        {/* Sign In Form Section */}
        <div className="p-8 space-y-6">
          {/* Logo */}
          <div className="flex items-center gap-2 mb-2">
            <img src={logo} alt="MyNneFarm Logo" className="h-10 w-10" />
            <h1 className="text-2xl font-bold text-[#2F5024]">My<span className="text-[#5E7E3F]">Nne</span>Farm</h1>
          </div>

          <h2 className="text-xl font-semibold text-[#4A2C1D]">Welcome Back</h2>
          <p className="text-sm text-[#6B4C3B]">Sign in to continue farming with us</p>

          <form className="space-y-4">
            <div className="relative">
              <Mail className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="email"
                placeholder="Email"
                className="w-full pl-10 pr-4 py-2 rounded-lg bg-white border border-gray-300 text-sm text-[#4A2C1D] focus:outline-none focus:ring-2 focus:ring-[#5E7E3F]"
                required
              />
            </div>

            <div className="relative">
              <Lock className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="password"
                placeholder="Password"
                className="w-full pl-10 pr-4 py-2 rounded-lg bg-white border border-gray-300 text-sm text-[#4A2C1D] focus:outline-none focus:ring-2 focus:ring-[#5E7E3F]"
                required
              />
            </div>

            <button
              type="submit"
              className="w-full bg-[#5E7E3F] text-white py-2 rounded-lg text-sm font-semibold hover:bg-[#466130] transition"
            >
              Sign In
            </button>
          </form>

          <p className="text-xs text-gray-600">
            Don’t have an account?{' '}
            <Link to="/auth/signup" className="text-[#2F5024] font-medium hover:underline">
              Sign Up
            </Link>
          </p>
        </div>

        {/* Visual Section */}
        <div className="bg-[#F3E7D5] p-6 flex flex-col justify-center items-center text-center">
          <h3 className="text-lg font-bold text-[#4A2C1D] mb-2">Grow with MyNneFarm</h3>
          <p className="text-sm text-[#5E7E3F] mb-4">
            Join our platform to build sustainable agriculture—secure your food future.
          </p>
          <img
            src="https://images.unsplash.com/photo-1568605114967-8130f3a36994?auto=format&fit=crop&w=600&q=60"
            alt="Farming Visual"
            className="rounded-lg shadow-md object-cover h-44 w-full"
          />
        </div>
      </div>
    </div>
  );
}
