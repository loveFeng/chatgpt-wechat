package gemini

import (
	"reflect"
	"testing"
)

func TestChatClient_CreateEmbedding(t *testing.T) {
	type fields struct {
		APIKey    string
		HttpProxy string
		Debug     bool
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    EmbeddingResponse
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				APIKey:    "AIxxxxxxx",
				HttpProxy: "http://127.0.0.1:7890",
				Debug:     true,
			},
			args: args{
				input: "Write a story about a magic backpack.",
			},
			want: EmbeddingResponse{
				Embedding: Embed{Values: []float64{0.008624583, -0.030451821, -0.042496547, -0.029230341, 0.05486475, 0.006694871, 0.004025645, -0.007294857, 0.0057651913, 0.037203953, 0.08070716, 0.032692064, 0.0015699493, -0.038671605, -0.021397846, 0.040436137, 0.040364444, 0.023915485, 0.03318194, -0.052099578, 0.007753789, -0.0028750803, -0.0038559572, -0.03839587, 0.031610277, -0.0024588231, 0.05350601, -0.035613116, -0.035775036, 0.045701347, -0.030365199, -0.014816799, -0.040846597, -0.014294212, 0.008432598, -0.07015665, -0.005973285, 0.020774437, -0.019995548, 0.027437009, -0.0143762855, 0.0071297227, -0.048812605, 0.0017134936, 0.016833002, -0.04341425, -0.01071614, 0.029540878, 0.00026989548, -0.07512045, -0.0063251033, 0.017243758, 0.0030855879, -0.03900979, 0.0062045115, -0.03762957, -0.0002221458, 0.0033970037, -0.018224807, 0.020233013, -0.009443185, 0.016834496, -0.039400727, 0.025765473, 0.0064459303, -0.0010064961, -0.023396038, 0.04714727, 0.04311917, 0.011308989, -0.013833369, -0.06827331, 0.023071568, -0.03515085, -0.06426478, -0.07674637, 0.011010596, 0.014995057, -0.009893141, 0.0226066, -0.023858562, -0.04174958, 0.00030446844, -0.029835863, -0.049982175, 0.030680457, -0.0037228062, 0.007982671, 0.015907364, 0.059540056, -0.0698364, 0.01905883, 0.026681246, -0.029017935, 0.009239862, 0.07437943, -0.018931432, -0.014418681, -0.015227716, -0.016991543, -0.020227646, -0.030113006, -0.036909197, 0.0491838, 0.03691079, 0.020114211, 0.020616315, 0.035417195, 0.017378854, 0.0017591371, -0.052360915, -0.007504276, -0.02162204, -0.04277857, -0.030450603, -0.008929546, 0.022382222, 0.028581386, 0.031293616, -0.017000198, 0.04805261, -0.030170312, 0.016913159, -0.0008443405, 0.017210385, 0.01790196, 0.025434153, 0.014020954, 0.0463916, 0.055676837, -0.014117397, -0.06040255, 0.033837322, -0.0008005907, -0.00060394837, 0.035327226, 0.036272198, -0.03526632, 0.008720279, -0.01767251, 0.030635742, 0.03079541, -0.011152445, 0.008129438, -0.004437317, 0.06261552, -0.011166501, -0.00792765, 0.0626778, -0.03808373, 0.0010393296, 0.0012560948, -0.05420512, -0.001696204, 0.0057959175, 0.021863215, -0.0057427636, -0.005779428, 0.009948935, -0.024309319, 0.03490945, 0.05541324, 0.010009066, -0.00690594, -0.017368019, -0.0020743837, 0.016718129, -0.021815343, 0.016868921, -0.016602708, -0.012883013, -0.049588937, -0.034187913, -0.034272812, -0.005009027, -0.06445695, 0.0061878716, -0.025500957, -0.0136196995, 0.009936822, -0.07557129, 0.0019269945, 0.007851136, -0.0005730017, 0.015097395, -0.02793086, 0.07649703, -0.011246095, -0.00988598, -0.0095420005, -0.010617724, -0.02795932, -0.0074260943, -0.0011066246, 0.030510733, 0.04752876, 0.0040175403, 0.029044962, 0.047818206, -0.018723032, -0.0415435, 0.0996901, 0.006733833, 0.026475549, 0.028504595, 0.039723564, 0.10685063, -0.09093502, -0.040105067, -0.010830562, -0.016954549, 0.040276904, -0.06309, 0.0122314235, 0.04197765, 0.021913808, 0.024538448, 0.03143963, 0.035233174, -0.049595617, 0.031046454, 0.012546503, -0.063403584, 0.029301276, 0.009593253, 0.08471234, -0.052641954, 0.06801721, -0.010078849, -0.03664156, -0.00001225098, 0.014980443, -0.015443251, -0.063587464, 0.0649348, 0.03656039, 0.00012944145, 0.04090392, -0.067475125, 0.042220943, -0.049328692, 0.00013846974, 0.030628476, -0.0044686855, -0.06414449, -0.0035188058, -0.021508386, 0.014263058, 0.0023899209, 0.0044664415, 0.011860193, -0.05595765, 0.03968002, 0.026143683, -0.04310548, 0.019457595, -0.036821175, -0.004706372, -0.008448093, 0.0095680095, 0.02663876, -0.017718185, 0.0521761, -0.05751985, -0.03382739, -0.00005254058, -0.007237099, -0.03678753, 0.0004373296, 0.068935804, 0.024607658, -0.07383697, 0.0745026, -0.020278804, -0.02233648, -0.043527547, -0.0005897141, -0.008819973, 0.05522694, -0.041430607, 0.01485464, 0.03093516, 0.027958557, -0.041524798, -0.04165515, -0.032893553, -0.03968652, -0.053652477, 0.017770097, 0.009334136, -0.05586768, -0.028391907, -0.032775786, -0.048513874, -0.053598277, 0.026337227, -0.016223265, 0.051107723, 0.043397397, -0.011614245, -0.051782615, -0.0044690934, 0.036513854, -0.059794012, 0.021193227, 0.022977995, -0.037308924, -0.04654618, 0.039977968, 0.0070000333, 0.010082792, -0.041809354, -0.06859667, 0.03696839, 0.08448864, 0.036238268, -0.040010847, 0.014791712, -0.071675524, 0.038495533, -0.025405306, 0.119683675, 0.053742535, -0.05001289, 0.013715115, 0.020359106, -0.011968625, 0.080088414, -0.036633175, 0.0514321, -0.092830576, -0.011293311, -0.011462946, -0.005365982, 0.0068834354, 0.0033007269, -0.061453447, -0.0018337568, -0.03999207, -0.0020025445, 0.030325854, -0.028261486, -0.0024511546, -0.04857929, -0.005050297, -0.013459029, -0.014253672, 0.03093196, 0.02680012, -0.023344921, 0.029151637, 0.06343295, -0.020851089, -0.013067708, -0.047613945, -0.019634524, 0.04799423, -0.0030165066, 0.023077987, -0.018307852, -0.02367432, 0.04621804, -0.00904888, -0.004921491, -0.011499991, -0.03138275, 0.00737706, -0.030905176, 0.0045861388, 0.022925997, -0.016103206, -0.037664305, -0.009711344, -0.041544404, -0.019569533, -0.039040513, -0.023987805, -0.020657333, -0.019713132, 0.012216924, -0.028459836, -0.007854262, 0.03432555, 0.018948609, 0.032789946, -0.002173598, 0.072268486, 0.044727862, -0.0047442573, 0.026857385, -0.004011348, -0.035373602, 0.064441904, 0.06910071, -0.011144723, -0.02612964, -0.00051150133, -0.058811516, 0.016943831, -0.013993827, -0.011681567, -0.0486106, -0.010806049, -0.009677699, -0.0075841006, -0.013452097, 0.050830264, 0.0069918637, -0.028301245, -0.0226844, 0.020452417, 0.038501225, 0.027227988, -0.09067933, -0.03149255, -0.02733588, 0.062468164, -0.011298025, 0.00020811577, 0.02480444, 0.030436065, -0.01722424, 0.015863098, 0.021556586, -0.035869934, -0.0105872825, -0.012277281, -0.050149817, 0.00007532577, 0.014090748, 0.0022058648, -0.0077205827, 0.01042793, -0.036767684, -0.019879367, -0.015746206, 0.017803842, 0.012614761, -0.00880104, -0.02583725, 0.021856116, -0.035151184, 0.0795235, 0.003733422, -0.042395752, -0.030227657, 0.017081745, -0.064787105, 0.047976263, -0.06614391, 0.046755534, -0.09351948, -0.017798718, -0.06981937, -0.048591003, -0.036941074, -0.0063392953, 0.0723561, -0.050979175, 0.024858551, 0.022146545, -0.04561866, -0.05629803, -0.03543026, 0.01992356, -0.02645938, 0.015476739, 0.006532406, 0.016006118, 0.021703305, -0.008074443, -0.013993359, 0.025270082, 0.054084614, -0.03723426, 0.00922647, -0.060977213, 0.022743328, 0.0005817427, -0.043921262, 0.0162521, -0.046245884, 0.02920244, 0.0137127, -0.0004419291, 0.0062954514, 0.0075316126, -0.018215746, -0.047283698, 0.06998149, -0.033327773, -0.0004236732, -0.0031994286, -0.007056563, -0.043460306, 0.0015354953, -0.01488144, -0.032937713, 0.009287482, 0.014544634, 0.034704477, -0.038788475, 0.0057188864, -0.041650325, 0.058672834, -0.037773453, 0.042793583, 0.068971485, -0.060984336, -0.003988655, -0.0028867219, 0.0067583215, -0.018067246, -0.0239257, 0.021824041, -0.002594604, 0.019783823, 0.010555229, 0.03585786, -0.054828122, 0.056835514, 0.0039436664, -0.029769812, 0.01487401, 0.018713957, -0.04180365, 0.065259494, -0.006946442, -0.008461352, -0.041328337, 0.016176524, 0.06900452, -0.08757591, -0.026511896, -0.021864926, -0.045825586, -0.0029127926, -0.036086105, 0.049907155, -0.03262437, 0.008395844, 0.014912004, 0.016121961, 0.038142838, -0.019255152, -0.032568473, 0.029633947, -0.05650531, 0.01703388, -0.0049108807, -0.033846553, -0.032649934, 0.034349475, -0.052442193, 0.035418052, -0.025731172, -0.028500304, -0.022009343, 0.0073188776, -0.02605774, -0.011230884, -0.016760005, -0.026268288, -0.030098971, 0.009599001, -0.012166129, -0.047288176, -0.0026035684, 0.046940323, 0.017147271, -0.03532738, -0.004257927, 0.023836099, -0.013437756, 0.038638394, -0.04540704, -0.0070548924, -0.000996806, -0.007153008, 0.03372742, 0.00090462615, 0.022542186, 0.056735456, 0.042577762, -0.034696132, 0.042536404, 0.021590313, 0.0077237147, 0.024994696, 0.029911542, -0.021255728, 0.030441552, -0.0483429, 0.04303822, 0.0286698, -0.0068607414, 0.036662962, -0.0063703014, -0.044340007, -0.031890824, 0.00036194356, -0.034090873, -0.00549679, 0.009660412, 0.042241063, 0.011368424, -0.004538653, -0.009493857, 0.0030975502, -0.0010478802, -0.020607537, 0.018744059, 0.015208846, -0.021333545, 0.03751383, 0.024116268, 0.07453785, -0.041588385, -0.03892425, -0.05235617, -0.040644005, 0.005042716, -0.020569988, -0.0129598, 0.13083012, -0.009011917, -0.00217832, 0.0077060633, 0.058262043, 0.015077671, 0.063272804, 0.1078087, 0.004448191, -0.053923953, -0.04362896, 0.09360521, 0.0066842767, -0.011016014, 0.044551995, 0.0015021093, -0.052759856, -0.009717925, 0.0034341498, 0.020852385, -0.0078668, 0.10094906, 0.07162882, -0.0748456, -0.027106045, 0.009101185, -0.029127726, -0.0017386917, -0.023493223, -0.027168266, -0.020215228, 0.00041417315, -0.033961166, -0.011669535, -0.0004906546, -0.012759002, -0.044284903, 0.04930086, 0.013013342, -0.020515632, 0.0126403915, 0.016976478, -0.08650424, -0.07489142, -0.04380144, 0.052320037, -0.06340725, 0.067897715, 0.031920537, -0.038168993, 0.036792386, 0.029663036, 0.022649394, 0.05061561, 0.00934687, 0.04729442, -0.018025605, 0.019651046, -0.0050999606, -0.0020830606, -0.007575653, 0.0045946045, 0.04751231, 0.007070753, -0.035760302, 0.018472316, 0.004339673, -0.06597283, -0.05489254, -0.011515522, 0.090681635, 0.007154289, 0.015031737, 0.008287731, 0.026016485, 0.0616728, -0.016931107, 0.018779512, -0.032710046, -0.010483889, 0.026504684, -0.020419342, -0.022554679, 0.025899567, 0.045513034, 0.00026808516, 0.03389962, -0.039920982, -0.0038337265, 0.0014569712, -0.009203633, -0.011793006, 0.014427106, 0.0086658755, -0.01721355, 0.08369377, 0.05515183, 0.03119344, 0.038981467, -0.034288254, -0.013515418, 0.06075744, -0.0258169, 0.034621883, 0.0012731912, -0.043584045, 0.04525766, -0.032612998, -0.020666298, 0.07351347, -0.050300013, 0.026697695, -0.0022883194, 0.0155193815, -0.017274313, -0.0020913866, -0.064670034, 0.018535795, -0.010191767, 0.08379303, 0.051132496, -0.057075754, 0.049261495, -0.011337851, -0.054149605, 0.03255013, -0.09124333, 0.03779213, 0.06664394, 0.00040837182, 0.028164629, -0.044449247, -0.012616811, 0.01718758, -0.013388284, 0.036616728, -0.009780496, 0.023196792, 0.0024103, 0.0152416425, -0.019779433, -0.014335527, 0.031857576, 0.012219593}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewChatClient(tt.fields.APIKey).WithDebug(tt.fields.Debug).
				WithHttpProxy(tt.fields.HttpProxy)
			got, err := c.CreateEmbedding(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEmbedding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmbedding() got = %v, want %v", got, tt.want)
			}
		})
	}
}
