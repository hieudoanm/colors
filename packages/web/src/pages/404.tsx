import { NextPage } from 'next';
import Link from 'next/link';

const NotFoundPage: NextPage = () => {
	return (
		<main className="flex min-h-screen flex-col items-center justify-center bg-neutral-900 px-4 text-neutral-100">
			<div className="flex flex-col items-center gap-6 text-center">
				<h1 className="text-6xl font-bold">404</h1>
				<p className="text-xl text-neutral-400">Sorry, the page you are looking for does not exist.</p>
				<Link
					href="/"
					className="rounded-xl bg-neutral-100 px-6 py-3 font-medium text-neutral-900 transition hover:bg-neutral-200">
					Go back home
				</Link>
			</div>
		</main>
	);
};

export default NotFoundPage;
