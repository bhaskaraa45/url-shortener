import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/error404.css'

interface NotFoundProps {}

const NotFound: React.FC<NotFoundProps> = () => {
  return (
    <div className="not-found">
      <h1>Oops! Page Not Found</h1>
      <p>
        The page you're looking for doesn't seem to exist. You may have mistyped
        the URL or the page may have been removed.
      </p>
      <Link to="/">Go Back Home</Link>
    </div>
  );
};

export default NotFound;
