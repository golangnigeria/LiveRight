import React from 'react';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <section className="bg-[#F3E7D5] text-[#1F3B17] min-h-screen">
      <div className="max-w-7xl mx-auto px-6 py-20">
        {/* Hero Section */}
        <div className="text-center mb-12">
          <h1 className="text-4xl md:text-5xl font-bold mb-4">
            Invest in Farms. Earn Food Credits.
          </h1>
          <p className="text-lg text-[#4A2C1D] max-w-xl mx-auto mb-6">
            MyNneFarm lets you grow your money by investing in real Nigerian farms. At harvest, receive food credits to redeem or cash out.
          </p>
          <Link
            to="/farms"
            className="bg-[#5E7E3F] mx-2 hover:bg-[#2F5024] text-white font-semibold py-3 px-6 rounded-lg transition-all"
          >
            Explore Farms
          </Link>
          <Link
            to="/farms/post"
            className="bg-[#5E7E3F] mx-2 hover:bg-[#2F5024] text-white font-semibold py-3 px-6 rounded-lg transition-all"
          >
            Post Farms
          </Link>
        </div>

        {/* Benefits Section */}
        <div className="grid md:grid-cols-3 gap-6 text-center mt-16">
          <div className="bg-white shadow-lg rounded-xl p-6">
            <h3 className="text-xl font-bold mb-2 text-[#4A2C1D]">Invest Smart</h3>
            <p className="text-[#4A2C1D]">Start with as low as ₦5,000 and support real farmers with real returns.</p>
          </div>
          <div className="bg-white shadow-lg rounded-xl p-6">
            <h3 className="text-xl font-bold mb-2 text-[#4A2C1D]">Earn Food Credits</h3>
            <p className="text-[#4A2C1D]">Get rewarded in food credits you can redeem for food or cash out.</p>
          </div>
          <div className="bg-white shadow-lg rounded-xl p-6">
            <h3 className="text-xl font-bold mb-2 text-[#4A2C1D]">Support Farmers</h3>
            <p className="text-[#4A2C1D]">Empower local agriculture and help reduce food insecurity in Nigeria.</p>
          </div>
        </div>

        {/* CTA Banner */}
        <div className="bg-[#1F3B17] text-white mt-20 rounded-lg p-10 text-center">
          <h2 className="text-2xl font-bold mb-3">Ready to grow with us?</h2>
          <p className="mb-5">Join hundreds of Nigerians securing their food future through smart farm investing.</p>
          <Link
            to="/farms"
            className="bg-[#A2B373] hover:bg-[#5E7E3F] text-[#1F3B17] font-bold py-2 px-5 rounded-lg transition-all"
          >
            View Available Farms
          </Link>
        </div>
      </div>
    </section>
  );
}

export default Home;
