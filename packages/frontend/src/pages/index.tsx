import { hex2hsl, hex2oklch, hex2rgb } from '@colors/utils/colors/code/hex';
import { getBrightness, randomHexColorCode } from '@colors/utils/colors/utils';
import { NextPage } from 'next';
import { useEffect, useState } from 'react';

const INITIAL_COLOR = '#171717';

export const HomePage: NextPage = () => {
  const [
    {
      hex = INITIAL_COLOR,
      hsl = hex2hsl(INITIAL_COLOR),
      rgb = hex2rgb(INITIAL_COLOR),
      oklch = hex2oklch(INITIAL_COLOR),
    },
    setState,
  ] = useState<{
    hex: string;
    hsl: { h: number; s: number; l: number } | null;
    rgb: { r: number; g: number; b: number } | null;
    oklch: { l: number; c: number; h: number } | null;
  }>({
    hex: INITIAL_COLOR,
    hsl: hex2hsl(INITIAL_COLOR),
    rgb: hex2rgb(INITIAL_COLOR),
    oklch: hex2oklch(INITIAL_COLOR),
  });

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === 'Space' || event.key === ' ') {
        const hex = randomHexColorCode();
        setState((previous) => ({
          ...previous,
          hex,
          hsl: hex2hsl(hex),
          rgb: hex2rgb(hex),
          oklch: hex2oklch(hex),
        }));
      }
    };

    window.addEventListener('keydown', handleKeyDown);

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, []);

  return (
    <div
      className="flex h-full items-center justify-center p-4 md:p-8"
      style={{
        backgroundColor: hex,
        color: getBrightness(hex) ? '#ffffff' : '#101828',
      }}>
      <div className="flex w-full max-w-md flex-col divide-y divide-dotted rounded-xl border border-dotted text-center">
        <input
          id="hex"
          name="hex"
          placeholder="HEX"
          className="w-full px-3 py-1 text-center text-sm uppercase focus:outline-none"
          value={hex}
          onChange={(event) => {
            const newHex = event.target.value;
            setState((previous) => ({
              ...previous,
              hex: newHex,
              hsl: hex2hsl(newHex),
              rgb: hex2rgb(newHex),
              oklch: hex2oklch(newHex),
            }));
          }}
        />
        <input
          id="rgb"
          name="rgb"
          placeholder="RGB"
          className="w-full px-3 py-1 text-center text-sm focus:outline-none"
          value={`rgb(${rgb?.r}, ${rgb?.g}, ${rgb?.b})`}
          readOnly
        />
        <input
          id="hsl"
          name="hsl"
          placeholder="HSL"
          className="w-full px-3 py-1 text-center text-sm focus:outline-none"
          value={`hsl(${hsl?.h}, ${hsl?.s}, ${hsl?.l}%)`}
          readOnly
        />
        <input
          id="oklch"
          name="oklch"
          placeholder="HSL"
          className="w-full px-3 py-1 text-center text-sm focus:outline-none"
          value={`oklch(${oklch?.l}, ${oklch?.c}, ${oklch?.h}%)`}
          readOnly
        />
      </div>
    </div>
  );
};

export default HomePage;
