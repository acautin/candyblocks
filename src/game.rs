use bevy::input::{keyboard::KeyboardInput, ElementState};

use crate::prelude::*;

pub struct CandyBlocksPlugin;

impl Plugin for CandyBlocksPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system(init_game)
            .add_system(print_keyboard_events)
            .add_system(gamepad_system)
            .add_system(print_gamepad_events);
    }
}

fn init_game(mut commands: Commands) {
    commands.spawn_bundle(OrthographicCameraBundle::new_2d());
    commands.spawn_bundle(SpriteBundle {
        sprite: Sprite {
            custom_size: Some(Vec2::new(100.0, 100.0)),
            color: Color::RED,
            ..Default::default()
        },
        ..Default::default()
    });
    info!("Game initialized!");
}

fn print_keyboard_events(mut keyboard_input_events: EventReader<KeyboardInput>) {
    for event in keyboard_input_events.iter() {
        info!("{:?}", event);
        if matches!(event.key_code, Some(KeyCode::Escape)) && event.state == ElementState::Released {
            info!("Escape detected, bye bye.");
            std::process::exit(0);
        }
    }
}

fn print_gamepad_events(mut gamepad_events: EventReader<GamepadEvent>) {
    for event in gamepad_events.iter() {
        match &event {
            GamepadEvent(gamepad, GamepadEventType::Connected) => {
                info!("{:?} Connected", gamepad);
            }
            GamepadEvent(gamepad, GamepadEventType::Disconnected) => {
                info!("{:?} Disconnected", gamepad);
            }
            GamepadEvent(gamepad, GamepadEventType::ButtonChanged(button_type, value)) => {
                info!("{:?} of {:?} is changed to {}", button_type, gamepad, value);
            }
            GamepadEvent(gamepad, GamepadEventType::AxisChanged(axis_type, value)) => {
                info!("{:?} of {:?} is changed to {}", axis_type, gamepad, value);
            }
        }
    }
}


fn gamepad_system(
    gamepads: Res<Gamepads>,
    button_inputs: Res<Input<GamepadButton>>,
    button_axes: Res<Axis<GamepadButton>>,
    axes: Res<Axis<GamepadAxis>>,
) {
    for gamepad in gamepads.iter().cloned() {
        if button_inputs.just_pressed(GamepadButton(gamepad, GamepadButtonType::South)) {
            info!("{:?} just pressed South", gamepad);
        } else if button_inputs.just_released(GamepadButton(gamepad, GamepadButtonType::South)) {
            info!("{:?} just released South", gamepad);
        }

        let right_trigger = button_axes
            .get(GamepadButton(gamepad, GamepadButtonType::RightTrigger2))
            .unwrap();
        if right_trigger.abs() > 0.01 {
            info!("{:?} RightTrigger2 value is {}", gamepad, right_trigger);
        }

        let left_stick_x = axes
            .get(GamepadAxis(gamepad, GamepadAxisType::LeftStickX))
            .unwrap();
        if left_stick_x.abs() > 0.01 {
            info!("{:?} LeftStickX value is {}", gamepad, left_stick_x);
        }

        let left_stick_pad_x = axes
            .get(GamepadAxis(gamepad, GamepadAxisType::DPadX))
            .unwrap();
        if left_stick_pad_x.abs() > 0.1 {
            info!("{:?} Dpadx value is {}", gamepad, left_stick_pad_x);
        }
    }
}

