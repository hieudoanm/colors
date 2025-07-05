import Link from 'next/link';
import { FC } from 'react';

export const CallToAction: FC = () => {
	return (
		<section className="w-full py-16">
			<div className="mx-auto flex max-w-3xl flex-col gap-y-8 text-center">
				<h3 className="text-2xl font-bold sm:text-3xl">Working with colors just got easier</h3>
				<p className="text-neutral-500">
					Explore, pick, and tweak perfect colors for your design or code — all in your browser, no sign-up needed.
				</p>
				<div>
					<Link href="/app">
						<button type="button" className="cursor-pointer rounded-full border border-neutral-800 px-6 py-3">
							Launch App
						</button>
					</Link>
				</div>
			</div>
		</section>
	);
};
