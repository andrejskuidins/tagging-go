local wezterm = require("wezterm")
local act = wezterm.action
local config = {}
config = {
    color_scheme = 'Tomorrow Night',
    hide_tab_bar_if_only_one_tab = true,
    font_size = 11.0,
    font = wezterm.font 'Liberation Mono',
    enable_wayland = true,
    scrollback_lines = 20000,
    enable_scroll_bar = true,
    disable_default_key_bindings = true,
    mouse_wheel_scrolls_tabs = true,
    hide_mouse_cursor_when_typing = false,
    initial_cols = 189,
    initial_rows = 80,
}
config.mouse_bindings = {
    {
        event = { Down = { streak = 1, button = "Right" } },
        mods = "NONE",
        action = wezterm.action_callback(function(window, pane)
            local has_selection = window:get_selection_text_for_pane(pane) ~= ""
            if has_selection then
                    window:perform_action(act({ PasteFrom = "PrimarySelection" }), pane)
            else
                    window:perform_action(act({ PasteFrom = "Clipboard" }), pane)
            end
        end),
    },
}
config.keys = {
    {
      key = 't',
      mods = 'CTRL',
      action = act.SpawnTab 'CurrentPaneDomain',
    },
}

return config
