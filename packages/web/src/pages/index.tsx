import { CallToAction } from '@colors/components/landing/CallToAction';
import { Features } from '@colors/components/landing/Features';
import { Footer } from '@colors/components/landing/Footer';
import { Hero } from '@colors/components/landing/Hero';
import { Navbar } from '@colors/components/shared/Navbar';
import { NextPage } from 'next';

const HomePage: NextPage = () => {
	return (
		<>
			<Navbar />
			<div className="w-full border-t border-neutral-800" />
			<Hero />
			<div className="w-full border-t border-neutral-800" />
			<Features />
			<div className="w-full border-t border-neutral-800" />
			<CallToAction />
			<div className="w-full border-t border-neutral-800" />
			<Footer />
		</>
	);
};

export default HomePage;
