import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface MenuState {
  openItem: string[];
  drawerOpen: boolean;
}

const initialState: MenuState = {
  openItem: ['dashboard'],
  drawerOpen: true,
};

const menuSlice = createSlice({
  name: 'menu',
  initialState,
  reducers: {
    activeItem(state, action: PayloadAction<any>) {
      state.openItem = action.payload;
    },
    openDrawer(state, action: PayloadAction<any>) {
      state.drawerOpen = action.payload;
    },
  },
});

export const { activeItem, openDrawer } = menuSlice.actions;

export default menuSlice.reducer;
