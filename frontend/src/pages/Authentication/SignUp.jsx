import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import logo from '../../assets/nnefarm.png';
import {
  User,
  Tractor,
  Leaf,
  Phone,
  Gift,
  Mail,
  Lock,
  CheckCircle
} from 'lucide-react';

export default function SignUp() {
  const [role, setRole] = useState('investor');
  const [acceptedTerms, setAcceptedTerms] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!acceptedTerms) {
      alert("Please accept the terms and conditions to proceed.");
      return;
    }

    // Proceed with API submission...
    alert(`Signed up successfully as ${role}`);
  };

  return (
    <div className="min-h-screen pt-24 bg-gradient-to-br from-[#FAF6EF] to-[#FFFFFF] px-4">
      <div className="max-w-5xl mx-auto bg-white rounded-2xl shadow-xl grid md:grid-cols-2 overflow-hidden">
        
        {/* Form Section */}
        <div className="p-8 md:p-10 space-y-6">
          {/* Logo */}
          <div className="flex items-center gap-3 mb-2">
            <img src={logo} alt="MyNneFarm Logo" className="h-10 w-10" />
            <h1 className="text-2xl font-bold text-[#2F5024]">
              My<span className="text-[#5E7E3F]">Nne</span>Farm
            </h1>
          </div>

          <h2 className="text-xl font-semibold text-[#4A2C1D]">Create an Account</h2>

          {/* Role Selector */}
          <div className="flex gap-4 text-sm font-medium text-[#1F3B17]">
            <button
              onClick={() => setRole('investor')}
              className={`flex items-center gap-2 px-4 py-2 rounded-lg border ${
                role === 'investor' ? 'bg-[#E9F1E1] border-[#5E7E3F]' : 'border-gray-300'
              }`}
            >
              <Leaf className="w-4 h-4" />
              Investor
            </button>
            <button
              onClick={() => setRole('farmer')}
              className={`flex items-center gap-2 px-4 py-2 rounded-lg border ${
                role === 'farmer' ? 'bg-[#E9F1E1] border-[#5E7E3F]' : 'border-gray-300'
              }`}
            >
              <Tractor className="w-4 h-4" />
              Farmer
            </button>
          </div>

          {/* Form */}
          <form className="space-y-4" onSubmit={handleSubmit}>
            {/* Full Name */}
            <div className="relative">
              <User className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="text"
                placeholder="Full Name"
                className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-[#5E7E3F] focus:outline-none"
                required
              />
            </div>

            {/* Email */}
            <div className="relative">
              <Mail className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="email"
                placeholder="Email"
                className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-[#5E7E3F] focus:outline-none"
                required
              />
            </div>

            {/* Phone */}
            <div className="relative">
              <Phone className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="tel"
                placeholder="Phone Number"
                className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-[#5E7E3F] focus:outline-none"
                required
              />
            </div>

            {/* Promo Code */}
            <div className="relative">
              <Gift className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="text"
                placeholder="Promo Code (if any)"
                className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-[#5E7E3F] focus:outline-none"
              />
            </div>

            {/* Password */}
            <div className="relative">
              <Lock className="absolute top-3 left-3 h-5 w-5 text-[#4A2C1D]" />
              <input
                type="password"
                placeholder="Password"
                className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-[#5E7E3F] focus:outline-none"
                required
              />
            </div>

            {/* Terms and Conditions */}
            <div className="flex items-start gap-2 text-sm">
              <input
                type="checkbox"
                id="terms"
                checked={acceptedTerms}
                onChange={() => setAcceptedTerms(!acceptedTerms)}
                className="mt-1"
              />
              <label htmlFor="terms" className="text-[#4A2C1D]">
                I accept the{' '}
                <Link
                  to="/terms"
                  className="text-[#2F5024] underline hover:text-[#1c3711] font-medium"
                >
                  Terms and Conditions
                </Link>
              </label>
            </div>

            {/* Submit */}
            <button
              type="submit"
              className="w-full py-3 bg-[#5E7E3F] hover:bg-[#4a6632] text-white font-semibold rounded-lg transition"
            >
              Sign Up as {role.charAt(0).toUpperCase() + role.slice(1)}
            </button>
          </form>

          <p className="text-xs text-gray-600">
            Already have an account?{' '}
            <Link to="/auth/signin" className="text-[#2F5024] font-medium hover:underline">
              Sign In
            </Link>
          </p>
        </div>

        {/* Ad Section */}
        <div className="bg-[#F3E7D5] px-6 py-10 flex flex-col justify-center items-center text-center">
          <h3 className="text-lg font-bold text-[#4A2C1D] mb-2">Join MyNneFarm Today</h3>
          <p className="text-sm text-[#5E7E3F] mb-4">
            Empowering farmers. Rewarding investors. Let’s grow together!
          </p>
          <img
            src="https://images.unsplash.com/photo-1607746882042-944635dfe10e?auto=format&fit=crop&w=600&q=60"
            alt="Farm Visual"
            className="rounded-xl shadow-md object-cover h-44 w-full"
          />
        </div>
      </div>
    </div>
  );
}
