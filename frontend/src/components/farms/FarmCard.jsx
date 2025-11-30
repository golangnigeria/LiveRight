// components/farms/FarmCard.jsx
import React from 'react';
import { Leaf, MapPin, TrendingUp, PackageCheck } from 'lucide-react';

export default function FarmCard({ farm }) {
  const defaultImage =
    'https://images.unsplash.com/photo-1580327331400-dc0c0be4a79c?auto=format&fit=crop&w=800&q=60';

  const unitsLeft = farm.units_available - (farm.units_sold || 0);
  const totalPotentialRevenue = farm.price_per_unit * farm.units_available;
  const expectedReturn = totalPotentialRevenue * farm.expected_roi;

  return (
    <div className="bg-white/40 backdrop-blur-md border border-white/30 rounded-2xl shadow-md p-4 flex flex-col md:flex-row gap-4 transition-transform hover:scale-[1.015] max-w-full md:max-w-[90%] md:mx-auto">
      <img
        src={farm.image_url || defaultImage}
        alt={farm.title}
        className="w-full md:w-1/2 h-48 md:h-auto object-cover rounded-xl"
      />

      <div className="flex flex-col justify-between flex-1">
        <div>
          <h2 className="text-2xl font-bold text-[#1F3B17] mb-2">{farm.title}</h2>
          <p className="text-gray-700 text-sm mb-3 line-clamp-3">{farm.description}</p>

          <div className="text-sm text-gray-800 flex flex-col gap-1 mb-2">
            <p className="flex items-center gap-2">
              <MapPin size={16} className="text-[#5E7E3F]" />
              {farm.location}
            </p>
            <p className="flex items-center gap-2">
              <Leaf size={16} className="text-[#5E7E3F]" />
              Crop: <span className="font-medium">{farm.crop}</span>
            </p>
            <p className="flex items-center gap-2">
              <TrendingUp size={16} className="text-[#5E7E3F]" />
              ROI: {Math.round(farm.expected_roi * 100 - 100)}%
            </p>
            <p className="flex items-center gap-2">
              <PackageCheck size={16} className="text-[#5E7E3F]" />
              Units Left: <span className="font-semibold text-[#1F3B17]">{unitsLeft}</span>
            </p>
          </div>

          <div className="w-full bg-gray-300/50 rounded-full h-3 mb-2">
            <div
              className="bg-[#5E7E3F] h-3 rounded-full transition-all duration-500"
              style={{
                width: `${(farm.units_sold / farm.units_available) * 100 || 0}%`,
              }}
            />
          </div>

          <p className="text-xs text-[#1F3B17] mb-1">
            Progress: {farm.units_sold || 0} / {farm.units_available} units sold
          </p>

          <p className="text-sm text-[#1F3B17] mb-2">
            <span className="font-semibold">Expected Harvest Return:</span>{' '}
            ₦{expectedReturn.toLocaleString()}
          </p>

          <p className="text-lg font-bold text-[#5E7E3F]">
            ₦{farm.price_per_unit.toLocaleString()} / unit
          </p>
        </div>

        <button
          onClick={() => alert(`Investing in: ${farm.title}`)}
          className="mt-4 bg-[#1F3B17] text-white px-4 py-2 rounded-xl hover:bg-[#345628] transition w-full md:w-max"
        >
          Invest Now
        </button>
      </div>
    </div>
  );
}
