// components/ui/LoadingSpinner.jsx
import React from "react";
import { Leaf } from "lucide-react";

export default function LoadingSpinner({ message = "Loading MyNneFarm..." }) {
  return (
    <div className="flex flex-col items-center justify-center min-h-[40vh] bg-gradient-to-br from-[#F3E7D5] to-[#ffffff] rounded-lg shadow-md p-6 w-full max-w-sm mx-auto">
      
      {/* Animated Spinner */}
      <div className="relative mb-4">
        <div className="h-12 w-12 rounded-full border-t-4 border-b-4 border-[#5E7E3F] animate-spin"></div>
        <Leaf className="absolute inset-0 m-auto text-[#5E7E3F] w-5 h-5" />
      </div>

      {/* Brand Name */}
      <h1 className="text-2xl font-extrabold text-[#1F3B17] tracking-wide mb-1">
        My<span className="text-[#5E7E3F]">Nne</span>Farm
      </h1>

      {/* Dynamic Message */}
      <p className="text-sm text-[#4A2C1D] font-medium animate-pulse mb-1">
        {message}
      </p>

      {/* Motto */}
      <p className="text-xs text-[#7E5D3F] italic">"Secure your food future"</p>
    </div>
  );
}
