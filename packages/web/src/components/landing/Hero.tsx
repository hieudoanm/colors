import { LinearText } from '@colors/components/shared/Linear';
import { APP_NAME } from '@colors/constants';
import Link from 'next/link';
import { FC } from 'react';

export const Hero: FC = () => {
	return (
		<section className="w-full py-20">
			<div className="mx-auto flex max-w-3xl flex-col gap-y-8 text-center">
				<h2 className="text-4xl font-extrabold text-neutral-100 sm:text-5xl">
					<LinearText>{APP_NAME}</LinearText>
				</h2>
				<p className="text-lg text-neutral-100">
					A fast, clean, and privacy-first color toolkit — pick, convert, and explore instantly in your browser.
				</p>
				<div>
					<Link href="/app">
						<button
							type="button"
							className="cursor-pointer rounded-full border border-neutral-700 bg-neutral-900 px-6 py-3">
							Launch App
						</button>
					</Link>
				</div>
			</div>
		</section>
	);
};
