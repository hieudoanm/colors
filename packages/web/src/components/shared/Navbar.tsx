import { APP_NAME } from '@colors/constants';
import Link from 'next/link';

export const Navbar = () => {
	return (
		<header className="w-full">
			<div className="container mx-auto flex items-center justify-between gap-x-2 px-4 py-2 md:px-8 md:py-4">
				<Link href="/">
					<h1 className="text-xl font-bold">{APP_NAME}</h1>
				</Link>
				<nav className="space-x-4 font-medium text-neutral-500">
					{[].map(({ id, href, label }) => {
						return (
							<Link key={id} href={href} className="text-sm md:text-base">
								{label}
							</Link>
						);
					})}
				</nav>
			</div>
		</header>
	);
};
