import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import {
  Menu,
  X,
  LogOut,
  Settings,
  BadgeCheck,
  BadgeX,
  Wallet,
  Leaf,
  ShoppingBag,
  Info,
  Tractor,
  User2,
} from 'lucide-react';
import logo from '../assets/nnefarm.png';

const Navbar = () => {
  const [mobileOpen, setMobileOpen] = useState(false);
  const [dropdownOpen, setDropdownOpen] = useState(false);

  // Simulated auth state (replace with your actual auth logic)
  const isAuthenticated = true;

  const user = isAuthenticated
    ? {
        name: 'Jane Doe',
        walletBalance: 58200,
        accountNumber: '0227458923',
        isVerified: true,
        avatarUrl: 'https://i.pravatar.cc/150?img=5',
      }
    : null;

  const formatCurrency = (amount) =>
    new Intl.NumberFormat('en-NG', {
      style: 'currency',
      currency: 'NGN',
      minimumFractionDigits: 0,
    }).format(amount);

  return (
    <nav className="bg-[#F3E7D5] text-[#4A2C1D] shadow-md fixed top-0 left-0 w-full z-50">
      <div className="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        {/* Logo */}
        <Link to="/" className="flex items-center gap-2">
          <img src={logo} alt="MyNneFarm Logo" className="h-10 w-10 object-contain" />
          <span className="text-xl font-bold tracking-tight text-[#1F3B17]">MyNneFarm</span>
        </Link>

        {/* Desktop Nav */}
        <div className="hidden md:flex gap-6 items-center text-sm font-medium">
          <Link to="/farms" className="hover:text-[#2F5024] flex items-center gap-1">
            <Tractor className="w-4 h-4" /> Farms
          </Link>
          <Link to="/marketplace" className="hover:text-[#2F5024] flex items-center gap-1">
            <ShoppingBag className="w-4 h-4" /> Marketplace
          </Link>
          <Link to="/about" className="hover:text-[#2F5024] flex items-center gap-1">
            <Info className="w-4 h-4" /> About
          </Link>
        </div>

        {/* User Dropdown */}
        <div className="hidden md:block relative">
          {isAuthenticated ? (
            <button
              onClick={() => setDropdownOpen(!dropdownOpen)}
              className="flex items-center gap-2 px-2 py-1 rounded-full hover:bg-[#E5D3B3] transition"
            >
              <img
                src={user.avatarUrl}
                alt="avatar"
                className="w-9 h-9 rounded-full border-2 border-[#5E7E3F] object-cover"
              />
              <span className="hidden md:inline text-sm font-semibold text-[#2F5024]">{user.name}</span>
            </button>
          ) : (
            <Link
              to="/auth/signin"
              className="bg-[#5E7E3F] text-white px-4 py-2 rounded-full text-sm font-semibold hover:bg-[#466130]"
            >
              Sign In
            </Link>
          )}

          {dropdownOpen && isAuthenticated && (
            <div className="absolute right-0 mt-2 w-64 bg-white border border-gray-200 rounded-lg shadow-lg z-50 p-4">
              <div className="flex items-center gap-3 mb-3">
                <img src={user.avatarUrl} alt="avatar" className="w-10 h-10 rounded-full object-cover" />
                <div>
                  <p className="font-semibold text-sm text-[#1F3B17]">{user.name}</p>
                  <p className="text-xs text-[#4A2C1D]">{user.accountNumber}</p>
                </div>
              </div>

              <div className="text-sm text-[#1F3B17] mb-2 space-y-1">
                <div className="flex items-center gap-2">
                  <Wallet className="w-4 h-4 text-[#2F5024]" />
                  <span>Wallet:</span>
                  <span className="font-semibold text-[#2F5024]">{formatCurrency(user.walletBalance)}</span>
                </div>
                <div className="flex items-center gap-2">
                  <User2 className="w-4 h-4 text-[#2F5024]" />
                  <span>Status:</span>
                  {user.isVerified ? (
                    <span className="flex items-center text-green-600 font-medium">
                      <BadgeCheck className="w-4 h-4" /> Verified
                    </span>
                  ) : (
                    <span className="flex items-center text-red-600 font-medium">
                      <BadgeX className="w-4 h-4" /> Not Verified
                    </span>
                  )}
                </div>
              </div>

              <hr className="my-2" />

              <div className="space-y-2">
                <Link to="/wallet" className="flex items-center gap-2 text-sm text-[#2F5024] hover:underline">
                  <Wallet className="w-4 h-4" /> Wallet
                </Link>
                <Link to="/settings" className="flex items-center gap-2 text-sm text-[#2F5024] hover:underline">
                  <Settings className="w-4 h-4" /> Settings
                </Link>
                <button className="flex items-center gap-2 text-sm text-red-600 hover:underline">
                  <LogOut className="w-4 h-4" /> Logout
                </button>
              </div>
            </div>
          )}
        </div>

        {/* Mobile Menu Toggle */}
        <div className="md:hidden">
          <button onClick={() => setMobileOpen(!mobileOpen)}>
            {mobileOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
          </button>
        </div>
      </div>

      {/* Mobile Menu */}
      {mobileOpen && (
        <div className="md:hidden px-4 py-4">
          <div className="bg-white rounded-xl shadow-lg p-4 space-y-4 border border-[#E5D3B3]">

            {user && (
              <div className="flex items-center gap-3 border-b pb-3">
                <img
                  src={user.avatarUrl}
                  alt="avatar"
                  className="w-10 h-10 rounded-full border border-[#5E7E3F]"
                />
                <div>
                  <p className="text-sm font-semibold text-[#1F3B17]">{user.name}</p>
                  <p className="text-xs text-[#4A2C1D]">{user.accountNumber}</p>
                </div>
              </div>
            )}

            <div className="flex flex-col gap-3 text-sm text-[#2F5024] font-medium">
              <Link to="/farms" onClick={() => setMobileOpen(false)} className="flex items-center gap-2">
                <Tractor className="w-4 h-4" /> Farms
              </Link>
              <Link to="/marketplace" onClick={() => setMobileOpen(false)} className="flex items-center gap-2">
                <ShoppingBag className="w-4 h-4" /> Marketplace
              </Link>
              <Link to="/about" onClick={() => setMobileOpen(false)} className="flex items-center gap-2">
                <Info className="w-4 h-4" /> About
              </Link>
            </div>

            <hr className="my-2 border-[#E5D3B3]" />

            <div className="flex flex-col gap-3 text-sm">
              {isAuthenticated ? (
                <>
                  <Link to="/wallet" onClick={() => setMobileOpen(false)} className="flex items-center gap-2 text-[#2F5024]">
                    <Wallet className="w-4 h-4" /> Wallet ({formatCurrency(user.walletBalance)})
                  </Link>
                  <Link to="/settings" onClick={() => setMobileOpen(false)} className="flex items-center gap-2 text-[#2F5024]">
                    <Settings className="w-4 h-4" /> Settings
                  </Link>
                  <button onClick={() => setMobileOpen(false)} className="flex items-center gap-2 text-red-600">
                    <LogOut className="w-4 h-4" /> Logout
                  </button>
                </>
              ) : (
                <Link
                  to="/signin"
                  onClick={() => setMobileOpen(false)}
                  className="bg-[#5E7E3F] text-white px-4 py-2 rounded-full text-center hover:bg-[#466130]"
                >
                  Sign In
                </Link>
              )}
            </div>
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;
