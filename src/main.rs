mod game;

mod prelude {
    pub use bevy::prelude::*;
    pub use crate::game::*;
}

use prelude::*;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugin(CandyBlocksPlugin)
        .run();
}
