import React, { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import { fetchFarms } from "../../api/farms";
import LoadingSpinner from "../../ui/LoadingSpinner";
import ErrorMessage from "../../ui/ErrorMessage";
import FarmCard from "../../components/farms/FarmCard";

function Farms() {
  const [showFarms, setShowFarms] = useState(false);

  const { data: farms, isLoading, isError } = useQuery({
    queryKey: ["farms"],
    queryFn: fetchFarms,
  });

  useEffect(() => {
    const timer = setTimeout(() => {
      setShowFarms(true);
    }, 5000); // 5-second delay before showing farm content
    return () => clearTimeout(timer);
  }, []);

  return (
    <section className="bg-[#F3E7D5] min-h-screen px-4 py-16">
      <h1 className="text-3xl font-bold text-center text-[#1F3B17] mb-10">
        Available Farms
      </h1>

      {/* Show spinner while loading or during delay */}
      {(isLoading || !showFarms) && (
        <LoadingSpinner message="Fetching farms..." />
      )}

      {/* Show error if query failed */}
      {isError && <ErrorMessage message="Failed to load farms. Please try again later." />}

      {/* Show farms after delay + successful fetch */}
      {showFarms && farms && farms.length > 0 && (
        <div className="flex flex-col gap-6">
          {farms.map((farm) => (
            <FarmCard key={farm.id} farm={farm} />
          ))}
        </div>
      )}

      {/* No farms found */}
      {showFarms && !isLoading && farms?.length === 0 && (
        <div className="text-center text-[#1F3B17]">
          No farms available at the moment. Please check back later.
        </div>
      )}

      <div className="text-center mt-10">
        <p className="text-[#4A2C1D] text-sm">
          Join us in supporting local farmers and securing your food future!
        </p>
      </div>
    </section>
  );
}

export default Farms;
