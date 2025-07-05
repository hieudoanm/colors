import { FC } from 'react';

export const Features: FC = () => {
	return (
		<section className="py-16">
			<div className="container mx-auto p-4 text-center md:p-8">
				<h3 className="text-3xl font-semibold sm:text-4xl">Why Use Our Color Tools?</h3>
				<p className="mx-auto mt-4 max-w-3xl text-neutral-500">
					Discover, edit, and manage colors effortlessly — all within your browser. No sign-up, no clutter, just color
					precision.
				</p>
				<div className="mt-12 grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
					<div className="rounded-xl border border-neutral-800 p-6 shadow-sm">
						<h4 className="text-lg font-semibold">🎨 Intuitive Interface</h4>
						<p className="mt-2 text-sm text-neutral-500">
							Pick, blend, and preview colors with ease using our sleek, responsive design.
						</p>
					</div>
					<div className="rounded-xl border border-neutral-800 p-6 shadow-sm">
						<h4 className="text-lg font-semibold">🔐 Private & Local</h4>
						<p className="mt-2 text-sm text-neutral-500">
							Everything runs in your browser — your palettes and selections stay private.
						</p>
					</div>
					<div className="rounded-xl border border-neutral-800 p-6 shadow-sm">
						<h4 className="text-lg font-semibold">🌈 Powerful Tools</h4>
						<p className="mt-2 text-sm text-neutral-500">
							From HEX to HSL, contrast checks to palette exports — all the tools you need, in one place.
						</p>
					</div>
				</div>
			</div>
		</section>
	);
};
