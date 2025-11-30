import React from 'react';
import logo from '../assets/nnefarm.png';
import team1 from '../assets/team1.jpg';
import team2 from '../assets/team2.jpg';
import team3 from '../assets/team3.jpg';
import team4 from '../assets/team4.jpg';

export default function About() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-[#FFFDF8] to-[#FAF6EF] pt-24 pb-16 px-4">
      <div className="max-w-6xl mx-auto bg-white/80 backdrop-blur-lg shadow-2xl rounded-2xl p-6 md:p-12 space-y-12">

        {/* Header */}
        <div className="flex items-center gap-4">
          <img src={logo} alt="MyNneFarm Logo" className="w-12 h-12" />
          <h1 className="text-3xl md:text-4xl font-bold text-[#2F5024]">
            About <span className="text-[#5E7E3F]">MyNneFarm</span>
          </h1>
        </div>

        {/* Intro */}
        <div className="text-lg text-[#4A2C1D] leading-relaxed">
          <p>
            <strong>MyNneFarm</strong> is a modern agro-finance platform that connects individuals with real farm opportunities. We empower Nigerian farmers through access to funding while offering investors a chance to grow wealth through sustainable agriculture.
          </p>
        </div>

        {/* Mission & Vision */}
        <div className="grid md:grid-cols-2 gap-6">
          <div className="bg-[#F9F5EC] rounded-lg p-4 shadow">
            <h2 className="text-xl font-semibold text-[#3B311F] mb-2">🌾 Our Mission</h2>
            <p className="text-[#5E7E3F]">
              To simplify farming investments and provide direct capital to farmers, ensuring both social impact and financial returns.
            </p>
          </div>
          <div className="bg-[#F9F5EC] rounded-lg p-4 shadow">
            <h2 className="text-xl font-semibold text-[#3B311F] mb-2">🌍 Our Vision</h2>
            <p className="text-[#5E7E3F]">
              To lead Africa’s agricultural transformation by making farming accessible, profitable, and tech-driven.
            </p>
          </div>
        </div>

        {/* Core Values */}
        <div>
          <h2 className="text-xl font-semibold text-[#3B311F] mb-4">💚 Our Core Values</h2>
          <ul className="space-y-2 pl-5 list-disc text-[#4A2C1D]">
            <li><strong>Transparency</strong>: We operate with openness and integrity.</li>
            <li><strong>Empowerment</strong>: We elevate both farmers and investors.</li>
            <li><strong>Innovation</strong>: We harness technology for agricultural growth.</li>
            <li><strong>Trust</strong>: We deliver reliable, secure partnerships.</li>
          </ul>
        </div>

        {/* Meet the Team */}
        <div>
          <h2 className="text-xl font-semibold text-[#3B311F] mb-6 text-center">👨‍🌾 Meet the MyNneFarm Team</h2>
          <div className="grid sm:grid-cols-2 md:grid-cols-4 gap-6">
            {[{ name: 'Prince Dimkpa', role: 'Founder & CEO', img: team1 },
              { name: 'Adaobi Nnadi', role: 'CTO & Head of Product', img: team2 },
              { name: 'Chinedu Okafor', role: 'Marketing Lead', img: team3 },
              { name: 'Blessing Ojo', role: 'Lead Agronomist', img: team4 }]
              .map((member, idx) => (
                <div key={idx} className="bg-white p-4 rounded-xl shadow-md text-center space-y-2">
                  <img src={member.img} alt={member.name} className="w-24 h-24 mx-auto rounded-full object-cover" />
                  <p className="font-semibold text-[#2F5024]">{member.name}</p>
                  <p className="text-sm text-[#5E7E3F]">{member.role}</p>
                </div>
              ))}
          </div>
        </div>

        {/* Call to Action */}
        <div className="text-center mt-12">
          <p className="text-[#4A2C1D] text-lg font-medium">
            Whether you’re a farmer or an investor, there's a place for you at MyNneFarm.
          </p>
          <a
            href="/auth/signup"
            className="inline-block mt-6 px-6 py-2 bg-[#5E7E3F] text-white rounded-lg font-semibold hover:bg-[#466130] transition"
          >
            Join the Movement
          </a>
        </div>
      </div>
    </div>
  );
}
