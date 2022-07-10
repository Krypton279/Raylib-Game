package main

import "github.com/gen2brain/raylib-go/raylib"
import "strconv"
import "math/rand"
import "math"

var height = 840
var width = 1440 

func map_algorithm(seed int,num int)([9]int){
	seeder := rand.NewSource(int64(seed))
	y1 := rand.New(seeder)
	x2 := y1.Intn(953243534) % num
	x3 := y1.Intn(953243534) % num
	x4 := y1.Intn(953243534) % num
	x5 := y1.Intn(953243534) % num
	x6 := y1.Intn(953243534) % num
	x7 := y1.Intn(953243534) % num
	x8 := y1.Intn(953243534) % num
	x1 := y1.Intn(953243534) % num
	x9 := y1.Intn(953243534) % num
	x := [...]int{x1,x2,x3,x4,x5,x6,x7,x8,x9}
	return x
}

func quad_origin_relative(x , y int)([2]int){
	x1 := x - (x % width)
	y1 := y - (y % height)
	xy := [...]int{x1,y1}
	return xy
}

func main() {
	x := 3
	y := 3
	zx := [...][2]int{{0 ,0},{0 ,0},{0 ,0},{0 ,0},{0 ,0},{0 ,0},{0 ,0},{0 ,0},{0 ,0}}
	s := float32(5.0)
	m := map_algorithm(rand.Intn(4000000),10)
	sd := map_algorithm(rand.Intn(4000000),10)
	for i := 0; i <= 8; i++ {
		m[i] = sd[i]
	}
	fc := quad_origin_relative(0 ,0)
	gh := quad_origin_relative(0 ,0)
	rectangle := rl.Rectangle{float32((width/2)-5),float32((height/2)-5), 10, 10}
	camera := new(rl.Camera2D)
	camera.Target = rl.Vector2{rectangle.X+float32(5),rectangle.Y+float32(5)}
	camera.Offset = rl.Vector2{float32(width/2),float32(height/2)}
	camera.Zoom = float32(1)
	camera.Rotation = float32(0)
	rl.InitWindow(int32(width),int32(height),"Test")
	rl.SetTargetFPS(60)
	
	// Loading Images and Textures
	player_image := rl.LoadImage("player.png")
	player_texture := rl.LoadTextureFromImage(player_image)
	wand_image := rl.LoadImage("starter_wand.png")
	wand_texture := rl.LoadTextureFromImage(wand_image)

	// Loading Images and Textures End
	// Game Loop Start
	for !rl.WindowShouldClose() {

		//    Mapper and Movement Code Begin
		rl.BeginDrawing()
		rl.BeginMode2D(*camera)
		rl.ClearBackground(rl.Black)
		if true {
			fc = quad_origin_relative(int(rectangle.X),int(rectangle.Y))
			sd = map_algorithm(rand.Intn(40000000),10)
			zx[0] = quad_origin_relative(int(rectangle.X),int(rectangle.Y))
			rl.DrawText(strconv.Itoa(m[0]),int32(zx[0][0]+(width/2)),int32(zx[0][1]+(height/2)),30,rl.RayWhite)
			zx[1] = quad_origin_relative(int(rectangle.X-float32(width)),int(rectangle.Y))
			rl.DrawText(strconv.Itoa(m[0]),int32(zx[1][0]+(width/2)),int32(zx[1][1]+(height/2)),30,rl.RayWhite)
			zx[2] = quad_origin_relative(int(rectangle.X-float32(width)),int(rectangle.Y-float32(height)))
			rl.DrawText(strconv.Itoa(m[1]),int32(zx[2][0]+(width/2)),int32(zx[2][1]+(height/2)),30,rl.RayWhite)
			zx[3] = quad_origin_relative(int(rectangle.X-float32(width)),int(rectangle.Y+float32(height)))
			rl.DrawText(strconv.Itoa(m[2]),int32(zx[3][0]+(width/2)),int32(zx[3][1]+(height/2)),30,rl.RayWhite)
			zx[4] = quad_origin_relative(int(rectangle.X+float32(width)),int(rectangle.Y+float32(height)))
			rl.DrawText(strconv.Itoa(m[3]),int32(zx[4][0]+(width/2)),int32(zx[4][1]+(height/2)),30,rl.RayWhite)
			zx[5] = quad_origin_relative(int(rectangle.X+float32(width)),int(rectangle.Y-float32(height)))
			rl.DrawText(strconv.Itoa(m[4]),int32(zx[5][0]+(width/2)),int32(zx[5][1]+(height/2)),30,rl.RayWhite)
			zx[6] = quad_origin_relative(int(rectangle.X+float32(width)),int(rectangle.Y))
			rl.DrawText(strconv.Itoa(m[5]),int32(zx[6][0]+(width/2)),int32(zx[6][1]+(height/2)),30,rl.RayWhite)
			zx[7] = quad_origin_relative(int(rectangle.X),int(rectangle.Y+float32(height)))
			rl.DrawText(strconv.Itoa(m[6]),int32(zx[7][0]+(width/2)),int32(zx[7][1]+(height/2)),30,rl.RayWhite)
			zx[8] = quad_origin_relative(int(rectangle.X),int(rectangle.Y-float32(height)))
			rl.DrawText(strconv.Itoa(m[7]),int32(zx[8][0]+(width/2)),int32(zx[8][1]+(height/2)),30,rl.RayWhite)
		}
		rl.DrawTexture(player_texture,int32(rectangle.X) - 25,int32(rectangle.Y) - 35,rl.White)
		x = int(rectangle.X)
		y = int(rectangle.Y)
		if rl.IsKeyDown(68){
			rectangle.X = rectangle.X + s
		}
		if rl.IsKeyDown(65){
			rectangle.X = rectangle.X - s
		}
		if rl.IsKeyDown(83){
			rectangle.Y = rectangle.Y + s
		}
		if rl.IsKeyDown(87){
			rectangle.Y = rectangle.Y - s
		}
		gh = quad_origin_relative(int(rectangle.X),int(rectangle.Y))
		if (fc[0] == gh[0]) && (fc[1] == gh[1]) {
		} else {
			for i := 0; i <= 8; i++{
				m[i] = sd[i]
			}
		}
		// Mapper and Movement Code End

		// Additional Stuff Start
		// rl.DrawRectangle(int32(x)+8,int32(y),25,5,rl.Green)
		rect := rl.Rectangle{float32(x+8),float32(y),25,6}
		rl.DrawTextureEx(wand_texture,rl.Vector2{rect.X,rect.Y-5},float32((-1)*(180/math.Pi)*(math.Atan2(float64((rl.GetMouseX()-int32((width/2)-5))),float64(rl.GetMouseY()-int32((height/2)-5))))+90),1,rl.White)
		// Additional Stuff End

		rl.EndMode2D()
		camera.Target = rl.Vector2{rectangle.X+float32(5),rectangle.Y+float32(5)} // Camera Fixed to Rect
		var text = "X: " + strconv.Itoa(int(rl.GetMouseX())) + " Y: " + strconv.Itoa(int(rl.GetMouseY())) + " FPS: " + strconv.Itoa(int(rl.GetFPS()))
		rl.DrawText(text,0 ,0 ,5 ,rl.RayWhite)
		rl.EndDrawing()
	}
	// Game Loop End
	// Removing Textures and Images from Memory
	rl.UnloadImage(player_image)
	rl.UnloadImage(wand_image)
	rl.UnloadTexture(player_texture)
	rl.UnloadTexture(wand_texture)
	// Removing Textures and Image End
	rl.CloseWindow()
}
