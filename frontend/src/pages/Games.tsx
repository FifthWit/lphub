import React from 'react';

import GameEntry from '../components/GameEntry';
import { Game } from '../types/Game';
import "../css/Maps.css"

interface GamesProps {
    games: Game[];
}

const Games: React.FC<GamesProps> = ({ games }) => {

    const _page_load = () => {
      const loaders = document.querySelectorAll(".loader");
      loaders.forEach((loader) => {
          (loader as HTMLElement).style.display = "none";
      });
    }

    React.useEffect(() => {
        document.querySelectorAll(".games-page-item-body").forEach((game, index) => {
            game.innerHTML = "";
        });
        _page_load();
    }, []);

    return (
        <div className='games-page'>
            <section>
                <div className='games-page-content'>
                    <div className='games-page-item-content'>
                        {games.map((game, index) => (
                            <GameEntry game={game} key={index} />
                        ))}
                    </div>
                </div>
            </section>
        </div>
    );
};

export default Games;