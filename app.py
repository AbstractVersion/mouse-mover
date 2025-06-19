import pyautogui
import time
import math

def move_mouse_pattern():
    """
    Moves mouse in a specific pattern to simulate activity
    """
    # Disable fail-safe (optional - remove if you want fail-safe)
    pyautogui.FAILSAFE = True
    
    # Get screen size
    screen_width, screen_height = pyautogui.size()
    
    # Set starting position (center of screen)
    start_x = screen_width // 2
    start_y = screen_height // 2
    
    # Move to starting position
    pyautogui.moveTo(start_x, start_y, duration=0.5)
    time.sleep(0.5)
    
    # Pattern parameters
    shaft_length = 200
    head_radius = 50
    movement_speed = 0.02
    
    # Draw the shaft (vertical line)
    for i in range(shaft_length):
        pyautogui.moveTo(start_x, start_y + i, duration=movement_speed)
    
    # Draw the head (semicircle at the top)
    head_center_y = start_y + shaft_length
    for angle in range(0, 181, 5):  # 0 to 180 degrees
        x = start_x + head_radius * math.cos(math.radians(angle))
        y = head_center_y + head_radius * math.sin(math.radians(angle))
        pyautogui.moveTo(x, y, duration=movement_speed)
    
    # Draw two circles on the sides (base)
    ball_radius = 30
    ball_offset = 40
    
    # Left circle
    left_center_x = start_x - ball_offset
    left_center_y = start_y - 20
    
    for angle in range(0, 361, 10):
        x = left_center_x + ball_radius * math.cos(math.radians(angle))
        y = left_center_y + ball_radius * math.sin(math.radians(angle))
        pyautogui.moveTo(x, y, duration=movement_speed)
    
    # Right circle
    right_center_x = start_x + ball_offset
    right_center_y = start_y - 20
    
    for angle in range(0, 361, 10):
        x = right_center_x + ball_radius * math.cos(math.radians(angle))
        y = right_center_y + ball_radius * math.sin(math.radians(angle))
        pyautogui.moveTo(x, y, duration=movement_speed)
    
    # Return to center
    pyautogui.moveTo(start_x, start_y, duration=0.5)

def keep_active(interval_minutes=5):
    """
    Continuously move mouse at specified intervals
    """
    print(f"Starting mouse activity simulation...")
    print(f"Mouse will move every {interval_minutes} minutes")
    print("Press Ctrl+C to stop")
    
    try:
        while True:
            print(f"Moving mouse at {time.strftime('%H:%M:%S')}")
            move_mouse_pattern()
            
            # Wait for specified interval
            time.sleep(interval_minutes * 60)
            
    except KeyboardInterrupt:
        print("\nStopping mouse movement script")

if __name__ == "__main__":
    # Install required library first: pip install pyautogui
    
    # Option 1: Run once
    # move_mouse_pattern()
    
    # Option 2: Run continuously (uncomment the line below)
    keep_active(interval_minutes=3)  # Move every 3 minutes